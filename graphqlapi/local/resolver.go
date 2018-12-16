/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
package local

import (
	local_get "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get"
	local_get_meta "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get_meta"
)

// Resolving local GraphQL queries
type Resolver interface {
	local_get.Resolver
	local_get_meta.Resolver
}