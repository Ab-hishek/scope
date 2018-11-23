package kubernetes

import (
	mayav1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	"github.com/weaveworks/scope/report"
)

// CStorVolumeReplica interface
type CStorVolumeReplica interface {
	Meta
	GetNode() report.Node
	GetCStorVolume() string
	GetCStorPool() string
}

// cStorVolume represents cStor Volume Replica CR
type cStorVolumeReplica struct {
	*mayav1alpha1.CStorVolumeReplica
	Meta
}

// NewCStorVolumeReplica returns fresh CStorVolumeReplica instance
func NewCStorVolumeReplica(p *mayav1alpha1.CStorVolumeReplica) CStorVolumeReplica {
	return &cStorVolumeReplica{CStorVolumeReplica: p, Meta: meta{p.ObjectMeta}}
}

// GetNode returns updated node with CStor Volume details
func (p *cStorVolumeReplica) GetNode() report.Node {
	var cStorPoolNodeID string
	latests := map[string]string{
		NodeType:   "CStor Volume Replica",
		APIVersion: p.APIVersion,
	}
	if p.GetCStorVolume() != "" {
		latests[CStorVolumeName] = p.GetCStorVolume()
	}
	if p.GetCStorPool() != "" {
		cStorPoolNodeID = report.MakeCStorPoolNodeID(p.GetCStorPool())
	}

	if p.GetCStorPool() != "" {
		latests[CStorPoolUID] = p.GetCStorPool()
	}

	return p.MetaNode(report.MakeCStorVolumeNodeID(p.UID())).
		WithLatests(latests).
		WithAdjacent(cStorPoolNodeID)
}

func (p *cStorVolumeReplica) GetCStorVolume() string {
	cStorVolumeName := p.Labels()["cstorvolume.openebs.io/name"]
	return cStorVolumeName
}

func (p *cStorVolumeReplica) GetCStorPool() string {
	cStorVolumeUID := p.Labels()["cstorpool.openebs.io/uid"]
	return cStorVolumeUID
}
