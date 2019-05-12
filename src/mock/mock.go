package mock

import "searchDemo/src/data"

var (
	MockStructMap = map[string]map[string]data.Field{
		"1": map[string]data.Field{
			"id": data.Field{Type: "string", NameWithCase: "ID", ValueMap: map[string][]interface{}{
				"t1": []interface{}{MockTickets[0]},
				"t2": []interface{}{MockTickets[1]},
			}},
			"url": data.Field{Type: "string", NameWithCase: "URL", ValueMap: map[string][]interface{}{
				"http://t1": []interface{}{MockTickets[0]},
				"http://t2": []interface{}{MockTickets[1]},
			}},
			"externalid": data.Field{Type: "string", NameWithCase: "ExternalID", ValueMap: map[string][]interface{}{
				"et1": []interface{}{MockTickets[0]},
				"et2": []interface{}{MockTickets[1]},
			}},
			"createdat": data.Field{Type: "string", NameWithCase: "CreatedAt", ValueMap: map[string][]interface{}{
				"2019-05-11t11:00:01": []interface{}{MockTickets[0]},
				"2019-05-11t11:00:02": []interface{}{MockTickets[1]},
			}},
			"type": data.Field{Type: "string", NameWithCase: "Type", ValueMap: map[string][]interface{}{
				"incident": []interface{}{MockTickets[0], MockTickets[1]},
			}},
			"subject": data.Field{Type: "string", NameWithCase: "Subject", ValueMap: map[string][]interface{}{
				"test1": []interface{}{MockTickets[0]},
				"test2": []interface{}{MockTickets[1]},
			}},
			"description": data.Field{Type: "string", NameWithCase: "Description", ValueMap: map[string][]interface{}{
				"test description": []interface{}{MockTickets[0]},
				"":                 []interface{}{MockTickets[1]},
			}},
			"priority": data.Field{Type: "string", NameWithCase: "Priority", ValueMap: map[string][]interface{}{
				"high": []interface{}{MockTickets[0], MockTickets[1]},
			}},
			"status": data.Field{Type: "string", NameWithCase: "Status", ValueMap: map[string][]interface{}{
				"pending": []interface{}{MockTickets[0], MockTickets[1]},
			}},
			"submitterid": data.Field{Type: "int", NameWithCase: "SubmitterID", ValueMap: map[string][]interface{}{
				"1": []interface{}{MockTickets[0]},
				"2": []interface{}{MockTickets[1]},
			}},
			"assigneeid": data.Field{Type: "int", NameWithCase: "AssigneeID", ValueMap: map[string][]interface{}{
				"1": []interface{}{MockTickets[1]},
				"2": []interface{}{MockTickets[0]},
			}},
			"organizationid": data.Field{Type: "int", NameWithCase: "OrganizationID", ValueMap: map[string][]interface{}{
				"1": []interface{}{MockTickets[0], MockTickets[1]},
			}},
			"tags": data.Field{Type: "[]string", NameWithCase: "Tags", ValueMap: map[string][]interface{}{
				"tag1.1": []interface{}{MockTickets[0]},
				"tag1.2": []interface{}{MockTickets[0]},
				"tag2.1": []interface{}{MockTickets[1]},
				"tag2.2": []interface{}{MockTickets[1]},
			}},
			"hasincidents": data.Field{Type: "bool", NameWithCase: "HasIncidents", ValueMap: map[string][]interface{}{
				"false": []interface{}{MockTickets[0], MockTickets[1]},
			}},
			"dueat": data.Field{Type: "string", NameWithCase: "DueAt", ValueMap: map[string][]interface{}{
				"2019-05-13t11:00:01": []interface{}{MockTickets[0]},
				"2019-05-13t11:00:02": []interface{}{MockTickets[1]},
			}},
			"via": data.Field{Type: "string", NameWithCase: "Via", ValueMap: map[string][]interface{}{
				"web": []interface{}{MockTickets[0], MockTickets[1]},
			}},
		},
		"2": map[string]data.Field{
			"id": data.Field{Type: "int", NameWithCase: "ID", ValueMap: map[string][]interface{}{
				"1": []interface{}{MockUsers[0]},
				"2": []interface{}{MockUsers[1]},
			}},
			"url": data.Field{Type: "string", NameWithCase: "URL", ValueMap: map[string][]interface{}{
				"http://u1": []interface{}{MockUsers[0]},
				"http://u2": []interface{}{MockUsers[1]},
			}},
			"externalid": data.Field{Type: "string", NameWithCase: "ExternalID", ValueMap: map[string][]interface{}{
				"u1": []interface{}{MockUsers[0]},
				"u2": []interface{}{MockUsers[1]},
			}},
			"name": data.Field{Type: "string", NameWithCase: "Name", ValueMap: map[string][]interface{}{
				"test testa": []interface{}{MockUsers[0]},
				"test testb": []interface{}{MockUsers[1]},
			}},
			"alias": data.Field{Type: "string", NameWithCase: "Alias", ValueMap: map[string][]interface{}{
				"user 1": []interface{}{MockUsers[0]},
				"user 2": []interface{}{MockUsers[1]},
			}},
			"createdat": data.Field{Type: "string", NameWithCase: "CreatedAt", ValueMap: map[string][]interface{}{
				"2019-04-11t11:00:01": []interface{}{MockUsers[0]},
				"2019-04-11t11:00:02": []interface{}{MockUsers[1]},
			}},
			"active": data.Field{Type: "bool", NameWithCase: "Active", ValueMap: map[string][]interface{}{
				"true": []interface{}{MockUsers[0], MockUsers[1]},
			}},
			"verified": data.Field{Type: "bool", NameWithCase: "Verified", ValueMap: map[string][]interface{}{
				"true": []interface{}{MockUsers[0], MockUsers[1]},
			}},
			"shared": data.Field{Type: "bool", NameWithCase: "Shared", ValueMap: map[string][]interface{}{
				"true": []interface{}{MockUsers[0], MockUsers[1]},
			}},
			"locale": data.Field{Type: "string", NameWithCase: "Locale", ValueMap: map[string][]interface{}{
				"en-au": []interface{}{MockUsers[0], MockUsers[1]},
			}},
			"timezone": data.Field{Type: "string", NameWithCase: "TimeZone", ValueMap: map[string][]interface{}{
				"australia": []interface{}{MockUsers[0], MockUsers[1]},
			}},
			"lastloginat": data.Field{Type: "string", NameWithCase: "LastLoginAt", ValueMap: map[string][]interface{}{
				"2019-05-12t11:00:01": []interface{}{MockUsers[0]},
				"2019-05-12t11:00:02": []interface{}{MockUsers[1]},
			}},
			"email": data.Field{Type: "string", NameWithCase: "Email", ValueMap: map[string][]interface{}{
				"user1@test.com": []interface{}{MockUsers[0]},
				"user2@test.com": []interface{}{MockUsers[1]},
			}},
			"phone": data.Field{Type: "string", NameWithCase: "Phone", ValueMap: map[string][]interface{}{
				"9991": []interface{}{MockUsers[0]},
				"9992": []interface{}{MockUsers[1]},
			}},
			"signature": data.Field{Type: "string", NameWithCase: "Signature", ValueMap: map[string][]interface{}{
				"":                []interface{}{MockUsers[0]},
				"user signature2": []interface{}{MockUsers[1]},
			}},
			"organizationid": data.Field{Type: "int", NameWithCase: "OrganizationID", ValueMap: map[string][]interface{}{
				"1": []interface{}{MockUsers[0], MockUsers[1]},
			}},
			"tags": data.Field{Type: "[]string", NameWithCase: "Tags", ValueMap: map[string][]interface{}{
				"utag1.1": []interface{}{MockUsers[0]},
				"utag1.2": []interface{}{MockUsers[0]},
				"utag2.1": []interface{}{MockUsers[1]},
				"utag2.2": []interface{}{MockUsers[1]},
			}},
			"suspended": data.Field{Type: "bool", NameWithCase: "Suspended", ValueMap: map[string][]interface{}{
				"false": []interface{}{MockUsers[0], MockUsers[1]},
			}},
			"role": data.Field{Type: "string", NameWithCase: "Role", ValueMap: map[string][]interface{}{
				"admin": []interface{}{MockUsers[0]},
				"user":  []interface{}{MockUsers[1]},
			}},
		},
		"3": map[string]data.Field{
			"id": data.Field{Type: "int", NameWithCase: "ID", ValueMap: map[string][]interface{}{
				"1": []interface{}{MockOrganizations[0]},
			}},
			"url": data.Field{Type: "string", NameWithCase: "URL", ValueMap: map[string][]interface{}{
				"http://org1": []interface{}{MockOrganizations[0]},
			}},
			"externalid": data.Field{Type: "string", NameWithCase: "ExternalID", ValueMap: map[string][]interface{}{
				"o1": []interface{}{MockOrganizations[0]},
			}},
			"name": data.Field{Type: "string", NameWithCase: "Name", ValueMap: map[string][]interface{}{
				"test org1": []interface{}{MockOrganizations[0]},
			}},
			"domainnames": data.Field{Type: "[]string", NameWithCase: "DomainNames", ValueMap: map[string][]interface{}{
				"org1.1.com": []interface{}{MockOrganizations[0]},
				"org1.2.com": []interface{}{MockOrganizations[0]},
			}},
			"createdat": data.Field{Type: "string", NameWithCase: "CreatedAt", ValueMap: map[string][]interface{}{
				"2019-05-01t11:00:00": []interface{}{MockOrganizations[0]},
			}},
			"details": data.Field{Type: "string", NameWithCase: "Details", ValueMap: map[string][]interface{}{
				"details1": []interface{}{MockOrganizations[0]},
			}},
			"sharedtickets": data.Field{Type: "bool", NameWithCase: "SharedTickets", ValueMap: map[string][]interface{}{
				"false": []interface{}{MockOrganizations[0]},
			}},
			"tags": data.Field{Type: "[]string", NameWithCase: "Tags", ValueMap: map[string][]interface{}{
				"otag1.1": []interface{}{MockOrganizations[0]},
				"otag1.2": []interface{}{MockOrganizations[0]},
			}},
		},
	}

	MockTickets = []*data.Ticket{
		&data.Ticket{
			ID:             "t1",
			URL:            "http://t1",
			ExternalID:     "et1",
			CreatedAt:      "2019-05-11T11:00:01",
			Type:           "incident",
			Subject:        "Test1",
			Description:    "Test description",
			Priority:       "high",
			Status:         "pending",
			SubmitterID:    1,
			AssigneeID:     2,
			OrganizationID: 1,
			Tags:           []string{"Tag1.1", "Tag1.2"},
			HasIncidents:   false,
			DueAt:          "2019-05-13T11:00:01",
			Via:            "web",
		},
		&data.Ticket{
			ID:             "t2",
			URL:            "http://t2",
			ExternalID:     "et2",
			CreatedAt:      "2019-05-11T11:00:02",
			Type:           "incident",
			Subject:        "test2",
			Description:    "",
			Priority:       "high",
			Status:         "pending",
			SubmitterID:    2,
			AssigneeID:     1,
			OrganizationID: 1,
			Tags:           []string{"Tag2.1", "Tag2.2"},
			HasIncidents:   false,
			DueAt:          "2019-05-13T11:00:02",
			Via:            "web",
		},
	}
	MockUsers = []*data.User{
		&data.User{
			ID:             1,
			URL:            "http://u1",
			ExternalID:     "u1",
			Name:           "Test TestA",
			Alias:          "user 1",
			CreatedAt:      "2019-04-11T11:00:01",
			Active:         true,
			Verified:       true,
			Shared:         true,
			Locale:         "en-AU",
			TimeZone:       "Australia",
			LastLoginAt:    "2019-05-12T11:00:01",
			Email:          "user1@test.com",
			Phone:          "9991",
			Signature:      "",
			OrganizationID: 1,
			Tags:           []string{"utag1.1", "utag1.2"},
			Suspended:      false,
			Role:           "admin",
		},
		&data.User{
			ID:             2,
			URL:            "http://u2",
			ExternalID:     "u2",
			Name:           "Test TestB",
			Alias:          "user 2",
			CreatedAt:      "2019-04-11T11:00:02",
			Active:         true,
			Verified:       true,
			Shared:         true,
			Locale:         "en-AU",
			TimeZone:       "Australia",
			LastLoginAt:    "2019-05-12T11:00:02",
			Email:          "user2@test.com",
			Phone:          "9992",
			Signature:      "user signature2",
			OrganizationID: 1,
			Tags:           []string{"utag2.1", "utag2.2"},
			Suspended:      false,
			Role:           "user",
		},
	}
	MockOrganizations = []*data.Organization{
		&data.Organization{
			ID:            1,
			URL:           "http://org1",
			ExternalID:    "o1",
			Name:          "test org1",
			DomainNames:   []string{"org1.1.com", "org1.2.com"},
			CreatedAt:     "2019-05-01T11:00:00",
			Details:       "details1",
			SharedTickets: false,
			Tags:          []string{"otag1.1", "otag1.2"},
		},
	}
)
