package artemis

import (
	"context"

	"github.com/vromero/activemq-artemis-operator/pkg/apis/vromero/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {

	switch o := event.Object.(type) {

	case *v1alpha1.ArtemisCluster:
		err := sdk.Create(createArtemisStatefulSet(o))
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("Failed to create statefulset pod : %v", err)
			return err
		}

		err = sdk.Create(createArtemisService(o))
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("Failed to create service : %v", err)
			return err
		}

	}
	return nil
}

// deploymentForMemcached returns a memcached Deployment object
func createArtemisStatefulSet(cr *v1alpha1.ArtemisCluster) *appsv1.StatefulSet {

	var variantTag string;
	if len(cr.Spec.Variant) > 0 {
		variantTag += "-" + cr.Spec.Variant
	}

	labels := map[string]string{
		"app": "busy-box",
	}
	replicas := cr.Spec.Size

	dep := &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "StatefulSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Image: "vromero/activemq-artemis:" + cr.Spec.Version + variantTag,
						Name:  "activemq-artemis",
						Ports: []v1.ContainerPort{
							{
								ContainerPort: 8161,
								Name:          "http",
							},
							{
								ContainerPort: 61616,
								Name:          "core",
							},
							{
								ContainerPort: 5445,
								Name:          "hornetq",
							},
							{
								ContainerPort: 5672,
								Name:          "amqp",
							},
							{
								ContainerPort: 1883,
								Name:          "mqtt",
							},
							{
								ContainerPort: 61613,
								Name:          "stomp",
							}},
					}},
				},
			},
		},
	}
	addOwnerRefToObject(dep, asOwner(cr))
	return dep
}

func createArtemisService(cr *v1alpha1.ArtemisCluster) *v1.Service {
	var variantTag string;
	if len(cr.Spec.Variant) > 0 {
		variantTag += "-" + cr.Spec.Variant
	}

	labels := map[string]string{
		"app": "busy-box",
	}

	dep := &v1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Selector: labels,
			Ports: []v1.ServicePort{
				{
					Port: 8161,
					Name: "http",
					TargetPort: intstr.IntOrString{
						IntVal: 8161,
					},
				},
				{
					Port: 61616,
					Name: "core",
					TargetPort: intstr.IntOrString{
						IntVal: 61616,
					},
				},
				{
					Port: 5445,
					Name: "hornetq",
					TargetPort: intstr.IntOrString{
						IntVal: 5445,
					},
				},
				{
					Port: 5672,
					Name: "amqp",
					TargetPort: intstr.IntOrString{
						IntVal: 5445,
					},
				},
				{
					Port: 1883,
					Name: "mqtt",
					TargetPort: intstr.IntOrString{
						IntVal: 1883,
					},
				},
				{
					Port: 61613,
					Name: "stomp",
					TargetPort: intstr.IntOrString{
						IntVal: 61613,
					},
				},
			},
		},
	}
	addOwnerRefToObject(dep, asOwner(cr))
	return dep
}

// deploymentForMemcached returns a memcached Deployment object
func createArtemisDeployment(cr *v1alpha1.ArtemisCluster) *appsv1.Deployment {
	labels := map[string]string{
		"app": "busy-box",
	}
	replicas := cr.Spec.Size

	dep := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Image:   "memcached:1.4.36-alpine",
						Name:    "memcached",
						Command: []string{"memcached", "-m=64", "-o", "modern", "-v"},
						Ports: []v1.ContainerPort{{
							ContainerPort: 11211,
							Name:          "memcached",
						}},
					}},
				},
			},
		},
	}
	addOwnerRefToObject(dep, asOwner(cr))
	return dep
}

func createArtemisPod(cr *v1alpha1.ArtemisCluster) *v1.Pod {
	labels := map[string]string{
		"app": "busy-box",
	}
	return &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "busy-box",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "ArtemisCluster",
				}),
			},
			Labels: labels,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}

// addOwnerRefToObject appends the desired OwnerReference to the object
func addOwnerRefToObject(obj metav1.Object, ownerRef metav1.OwnerReference) {
	obj.SetOwnerReferences(append(obj.GetOwnerReferences(), ownerRef))
}

// asOwner returns an OwnerReference set as the artemis CR
func asOwner(m *v1alpha1.ArtemisCluster) metav1.OwnerReference {
	trueVar := true
	return metav1.OwnerReference{
		APIVersion: m.APIVersion,
		Kind:       m.Kind,
		Name:       m.Name,
		UID:        m.UID,
		Controller: &trueVar,
	}
}
