package scratchlint

import (
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func f(log logr.Logger, t *metav1.Time) {
	log.Info("msg", "notAfter", t)                                                      // want finding: pair 1
	log.Info("msg", "a", 1, "notAfter", t)                                              // want finding: pair 2
	log.V(4).Info("msg", "notAfter", t)                                                 // want finding: chained V()
	log.Error(nil, "msg", "notAfter", t)                                                // want finding: Error
	l2 := log.WithValues("notAfter", t)                                                 // want finding: WithValues
	log.Info("updating status fields", "notAfter", t, "notBefore", t, "renewalTime", t) // want finding: issue 6799 shape, pointer at pair 3
	log.Info("msg", "a", 1, "b", 2, "c", 3, "d", 4, "e", 5, "notAfter", t)              // want finding: pair 6
	l2.Info("msg", "notAfter", klog.SafePtr(t))                                         // no finding: wrapped
	log.Info("msg", "count", 42, "name", "x")                                           // no finding: no pointers
	log.Info("msg", "notAfter", t.String())                                             // no finding: explicit String call
	log.Error(nil, "msg", "err", error(nil))                                            // no finding: error values are not Stringers
}
