package kas

import (
	hyperv1 "github.com/openshift/hypershift/api/types/hypershift/v1beta1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func ReconcilePodDisruptionBudget(pdb *policyv1.PodDisruptionBudget, p *KubeAPIServerParams) error {
	if pdb.CreationTimestamp.IsZero() {
		pdb.Spec.Selector = &metav1.LabelSelector{
			MatchLabels: kasLabels(),
		}
	}

	p.OwnerRef.ApplyTo(pdb)

	var minAvailable int
	switch p.Availability {
	case hyperv1.SingleReplica:
		minAvailable = 0
	case hyperv1.HighlyAvailable:
		minAvailable = 1
	}
	pdb.Spec.MinAvailable = &intstr.IntOrString{Type: intstr.Int, IntVal: int32(minAvailable)}

	return nil
}
