package client

import (
	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	kafka "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ckafka/v20190819"
	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
	cmq "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cmq/v20190304"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	cynosdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cynosdb/v20190107"
	dc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dc/v20180410"
	dcdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dcdb/v20180411"
	es "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es/v20180416"
	lh "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
	mariadb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mariadb/v20170312"
	memcached "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/memcached/v20190318"
	mongodb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20190725"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	pg "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/postgres/v20170312"
	redis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
	sqlserver "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver/v20180328"
	tdmq "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdmq/v20200217"
	tse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tse/v20201207"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
	"net/http"
	"net/url"

	cos "github.com/tencentyun/cos-go-sdk-v5"

	"github.com/tencentyun/tencentcloud-exporter/pkg/config"
)

func NewMonitorClient(conf *config.TencentConfig) (*monitor.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "monitor.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "monitor.tencentcloudapi.com"
	}
	return monitor.NewClient(credential, conf.Credential.Region, cpf)
}

func NewMongodbClient(conf *config.TencentConfig) (*mongodb.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "mongodb.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "mongodb.tencentcloudapi.com"
	}
	return mongodb.NewClient(credential, conf.Credential.Region, cpf)
}

func NewCdbClient(conf *config.TencentConfig) (*cdb.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "cdb.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "cdb.tencentcloudapi.com"
	}
	return cdb.NewClient(credential, conf.Credential.Region, cpf)
}

func NewCvmClient(conf *config.TencentConfig) (*cvm.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "cvm.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	}
	return cvm.NewClient(credential, conf.Credential.Region, cpf)
}

func NewRedisClient(conf *config.TencentConfig) (*redis.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "redis.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "redis.tencentcloudapi.com"
	}
	return redis.NewClient(credential, conf.Credential.Region, cpf)
}

func NewDcClient(conf *config.TencentConfig) (*dc.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "dc.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "dc.tencentcloudapi.com"
	}
	return dc.NewClient(credential, conf.Credential.Region, cpf)
}

func NewClbClient(conf *config.TencentConfig) (*clb.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "clb.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "clb.tencentcloudapi.com"
	}
	return clb.NewClient(credential, conf.Credential.Region, cpf)
}

func NewVpvClient(conf *config.TencentConfig) (*vpc.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "vpc.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "vpc.tencentcloudapi.com"
	}
	return vpc.NewClient(credential, conf.Credential.Region, cpf)
}

func NewCbsClient(conf *config.TencentConfig) (*cbs.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "cbs.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "cbs.tencentcloudapi.com"
	}
	return cbs.NewClient(credential, conf.Credential.Region, cpf)
}

func NewSqlServerClient(conf *config.TencentConfig) (*sqlserver.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "sqlserver.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "sqlserver.tencentcloudapi.com"
	}
	return sqlserver.NewClient(credential, conf.Credential.Region, cpf)
}

func NewMariaDBClient(conf *config.TencentConfig) (*mariadb.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "mariadb.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "mariadb.tencentcloudapi.com"
	}
	return mariadb.NewClient(credential, conf.Credential.Region, cpf)
}

func NewESClient(conf *config.TencentConfig) (*es.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "es.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "es.tencentcloudapi.com"
	}
	return es.NewClient(credential, conf.Credential.Region, cpf)
}

func NewCMQClient(conf *config.TencentConfig) (*cmq.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "cmq.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "cmq.tencentcloudapi.com"
	}
	return cmq.NewClient(credential, conf.Credential.Region, cpf)
}

func NewPGClient(conf *config.TencentConfig) (*pg.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "postgres.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "postgres.tencentcloudapi.com"
	}
	return pg.NewClient(credential, conf.Credential.Region, cpf)
}

func NewMemcacheClient(conf *config.TencentConfig) (*memcached.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "memcached.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "memcached.tencentcloudapi.com"
	}
	return memcached.NewClient(credential, conf.Credential.Region, cpf)
}

func NewLighthouseClient(conf *config.TencentConfig) (*lh.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "lighthouse.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	}
	return lh.NewClient(credential, conf.Credential.Region, cpf)
}

func NewKafkaClient(conf *config.TencentConfig) (*kafka.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "ckafka.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "ckafka.tencentcloudapi.com"
	}
	return kafka.NewClient(credential, conf.Credential.Region, cpf)
}

func NewDCDBClient(conf *config.TencentConfig) (*dcdb.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "dcdb.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "dcdb.tencentcloudapi.com"
	}
	return dcdb.NewClient(credential, conf.Credential.Region, cpf)
}

func NewTDMQClient(conf *config.TencentConfig) (*tdmq.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "tdmq.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "tdmq.tencentcloudapi.com"
	}
	return tdmq.NewClient(credential, conf.Credential.Region, cpf)
}

func NewTseClient(conf *config.TencentConfig) (*tse.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "tse.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "tse.tencentcloudapi.com"
	}
	return tse.NewClient(credential, conf.Credential.Region, cpf)
}

func NewCynosdbClient(conf *config.TencentConfig) (*cynosdb.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "cynosdb.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "cynosdb.tencentcloudapi.com"
	}
	return cynosdb.NewClient(credential, conf.Credential.Region, cpf)
}

func NewCdnClient(conf *config.TencentConfig) (*cdn.Client, error) {
	credential := common.NewCredential(
		conf.Credential.AccessKey,
		conf.Credential.SecretKey,
	)
	cpf := profile.NewClientProfile()
	if conf.Credential.IsInternal == true {
		cpf.HttpProfile.Endpoint = "cdn.internal.tencentcloudapi.com"
	} else {
		cpf.HttpProfile.Endpoint = "cdn.tencentcloudapi.com"
	}
	return cdn.NewClient(credential, "", cpf)
}

func NewCosClient(conf *config.TencentConfig) (*cos.Client, error) {
	// 用于Get Service 查询, service域名暂时只支持外网
	su, _ := url.Parse("http://cos." + conf.Credential.Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: nil, ServiceURL: su}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.Credential.AccessKey,
			SecretKey: conf.Credential.SecretKey,
		},
	})
	return client, nil
}
