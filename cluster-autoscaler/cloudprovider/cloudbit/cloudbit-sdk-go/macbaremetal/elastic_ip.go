package macbaremetal

import (
	"context"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/cloudbit/cloudbit-sdk-go/common"

	cloudbitgo "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/cloudbit/cloudbit-sdk-go"
)

type ElasticIPAttachment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ElasticIP struct {
	ID         int                 `json:"id"`
	Product    common.BriefProduct `json:"product"`
	Location   common.Location     `json:"location"`
	Price      float64             `json:"price"`
	PublicIP   string              `json:"public_ip"`
	PrivateIP  string              `json:"private_ip"`
	Attachment ElasticIPAttachment `json:"attached_device"`
}

type ElasticIPList struct {
	Items      []ElasticIP
	Pagination cloudbitgo.Pagination
}

type ElasticIPCreate struct {
	LocationID int `json:"location_id,omitempty"`
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

type ElasticIPAttach struct {
	ElasticIPID        int `json:"elastic_ip_id"`
	NetworkInterfaceID int `json:"network_interface_id"`
}

type AttachedElasticIPService struct {
	client   cloudbitgo.Client
	deviceID int
}

func NewAttachedElasticIPService(client cloudbitgo.Client, deviceID int) AttachedElasticIPService {
	return AttachedElasticIPService{client: client, deviceID: deviceID}
}

func (a AttachedElasticIPService) List(ctx context.Context, cursor cloudbitgo.Cursor) (list ElasticIPList, err error) {
	list.Pagination, err = a.client.List(ctx, getAttachedElasticIPsPath(a.deviceID), cursor, &list.Items)
	return
}

func (a AttachedElasticIPService) Attach(ctx context.Context, body ElasticIPAttach) (elasticIP ElasticIP, err error) {
	err = a.client.Create(ctx, getAttachedElasticIPsPath(a.deviceID), body, &elasticIP)
	return
}

func (a AttachedElasticIPService) Detach(ctx context.Context, id int) (err error) {
	err = a.client.Delete(ctx, getSpecificAttachedElasticIPPath(a.deviceID, id))
	return
}

const (
	elasticIPsSegment         = "/v4/macbaremetal/elastic-ips"
	attachedElasticIPsSegment = "elastic-ips"
)

func getElasticIPsPath() string {
	return elasticIPsSegment
}

func getSpecificElasticIPPath(id int) string {
	return cloudbitgo.Join(elasticIPsSegment, id)
}

func getAttachedElasticIPsPath(deviceID int) string {
	return cloudbitgo.Join(getSpecificDevicePath(deviceID), attachedElasticIPsSegment)
}

func getSpecificAttachedElasticIPPath(deviceID, elasticIPID int) string {
	return cloudbitgo.Join(getAttachedElasticIPsPath(deviceID), elasticIPID)
}
