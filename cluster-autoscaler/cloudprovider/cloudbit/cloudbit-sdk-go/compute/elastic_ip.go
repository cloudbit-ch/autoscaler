package compute

import (
	"context"
	cloudbitgo "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/cloudbit/cloudbit-sdk-go"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/cloudbit/cloudbit-sdk-go/common"
)

type ElasticIPProduct = common.BriefProduct

type ElasticIPAttachment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ElasticIP struct {
	ID         int                 `json:"id"`
	Product    ElasticIPProduct    `json:"product"`
	Location   common.Location     `json:"location"`
	Price      float64             `json:"price"`
	PublicIP   string              `json:"public_ip"`
	PrivateIP  string              `json:"private_ip"`
	Attachment ElasticIPAttachment `json:"attached_instance"`
}

type ElasticIPList struct {
	Items      []ElasticIP
	Pagination cloudbitgo.Pagination
}

type ElasticIPCreate struct {
	LocationID int `json:"location_id"`
}

type ElasticIPService struct {
	client cloudbitgo.Client
}

func NewElasticIPService(client cloudbitgo.Client) ElasticIPService {
	return ElasticIPService{client: client}
}

func (e ElasticIPService) List(ctx context.Context, cursor cloudbitgo.Cursor) (list ElasticIPList, err error) {
	list.Pagination, err = e.client.List(ctx, getElasticIPsPath(), cursor, &list.Items)
	return
}

func (e ElasticIPService) Create(ctx context.Context, body ElasticIPCreate) (elasticIP ElasticIP, err error) {
	err = e.client.Create(ctx, getElasticIPsPath(), body, &elasticIP)
	return
}

func (e ElasticIPService) Delete(ctx context.Context, id int) (err error) {
	err = e.client.Delete(ctx, getSpecificElasticIPPath(id))
	return
}

const elasticIPsSegment = "/v4/compute/elastic-ips"

func getElasticIPsPath() string {
	return elasticIPsSegment
}

func getSpecificElasticIPPath(elasticIPID int) string {
	return cloudbitgo.Join(elasticIPsSegment, elasticIPID)
}
