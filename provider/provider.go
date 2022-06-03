// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package provider

import (
	"encoding/json"
	"fmt"

	"github.com/gocql/gocql"
)

func init() {
	unregisteredFactories := map[Type]Factory{
		LocalOnline:      localOnlineStoreFactory,
		RedisOnline:      redisOnlineStoreFactory,
<<<<<<< HEAD:provider/provider.go
		CassandraOnline:  cassandraOnlineStoreFactory,
=======
		DynamoDBOnline:   dynamodbOnlineStoreFactory,
>>>>>>> e8846b1 (Passed all Dyanamodb tests):ff/serving/provider/provider.go
		MemoryOffline:    memoryOfflineStoreFactory,
		PostgresOffline:  postgresOfflineStoreFactory,
		SnowflakeOffline: snowflakeOfflineStoreFactory,
		RedshiftOffline:  redshiftOfflineStoreFactory,
	}
	for name, factory := range unregisteredFactories {
		if err := RegisterFactory(name, factory); err != nil {
			panic(err)
		}
	}
}

type SerializedConfig []byte

type SerializedTableSchema []byte

type RedisConfig struct {
	Prefix   string
	Addr     string
	Password string
	DB       int
}

type DynamodbConfig struct {
	Prefix string
	Addr   string
	Region string
}

func (r RedisConfig) Serialized() SerializedConfig {
	config, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return config
}

func (r *RedisConfig) Deserialize(config SerializedConfig) error {
	err := json.Unmarshal(config, r)
	if err != nil {
		return err
	}
	return nil
}

type CassandraConfig struct {
	keyspace    string
	Addr        string
	session     *gocql.Session
	Consistency gocql.Consistency
}

func (r CassandraConfig) Serialized() SerializedConfig {
	config, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return config
}

func (r DynamodbConfig) Serialized() SerializedConfig {
	config, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return config
}

func (r *CassandraConfig) Deserialize(config SerializedConfig) error {
	err := json.Unmarshal(config, r)
	if err != nil {
		return err
	}
	return nil
}

func (r *DynamodbConfig) Deserialize(config SerializedConfig) error {
	err := json.Unmarshal(config, r)
	if err != nil {
		return err
	}
	return nil
}

type Provider interface {
	AsOnlineStore() (OnlineStore, error)
	AsOfflineStore() (OfflineStore, error)
	Type() Type
	Config() SerializedConfig
}

type BaseProvider struct {
	ProviderType   Type
	ProviderConfig SerializedConfig
}

func (provider BaseProvider) AsOnlineStore() (OnlineStore, error) {
	return nil, fmt.Errorf("%T cannot be used as an OnlineStore", provider)
}

func (provider BaseProvider) AsOfflineStore() (OfflineStore, error) {
	return nil, fmt.Errorf("%T cannot be used as an OfflineStore", provider)
}

func (provider BaseProvider) Type() Type {
	return provider.ProviderType
}

func (provider BaseProvider) Config() SerializedConfig {
	return provider.ProviderConfig
}

type Factory func(SerializedConfig) (Provider, error)

type Type string

var factories map[Type]Factory = make(map[Type]Factory)

func RegisterFactory(t Type, f Factory) error {
	if _, has := factories[t]; has {
		return fmt.Errorf("%s provider factory already exists", t)
	}
	factories[t] = f
	return nil
}

func Get(t Type, config SerializedConfig) (Provider, error) {
	f, has := factories[t]
	if !has {
		return nil, fmt.Errorf("no provider of type: %s", t)
	}
	return f(config)
}
