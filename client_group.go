package uim

import cloudevents "github.com/cloudevents/sdk-go/v2"

type GroupMemberRequest struct {
	Member *GroupMember
}

func (client *Client) AddGroupMember(member *GroupMember) error {
	ce := cloudevents.NewEvent()
	ce.SetSource(UIMProviderEventSource)
	ce.SetType(AddGroupMemberRequest)
	ce.SetData(cloudevents.ApplicationJSON, GroupMemberRequest{
		Member: member,
	})

	return client.SendEvent(&ce)
}
