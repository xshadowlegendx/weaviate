/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */
package database

import (
	graphql_local_get "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get"
	graphql_local_get_meta "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get_meta"
)

type dbClosingResolver struct {
	connectorLock ConnectorLock
}

func (dbcr *dbClosingResolver) Close() {
	dbcr.connectorLock.Unlock()
}

func (dbcr *dbClosingResolver) LocalGetClass(info *graphql_local_get.LocalGetClassParams) (interface{}, error) {
	connector := dbcr.connectorLock.Connector()
	thunk, err := connector.LocalGetClass(info)
	return thunk, err
}

func (dbcr *dbClosingResolver) LocalGetMeta(info *graphql_local_get_meta.LocalGetMetaParams) (interface{}, error) {
	connector := dbcr.connectorLock.Connector()
	return connector.LocalGetMeta(info)
}
