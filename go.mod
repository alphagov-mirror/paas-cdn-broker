module github.com/18F/cf-cdn-service-broker

go 1.12

require (
	code.cloudfoundry.org/lager v0.0.0-20170223024724-de8e9c6c6e47
	github.com/aws/aws-sdk-go v1.6.27
	github.com/cenkalti/backoff v2.1.1+incompatible // indirect
	github.com/cloudfoundry-community/go-cfclient v0.0.0-20170502130034-23156ed0e73b
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20190515213511-eb9f6a1743f3 // indirect
	github.com/drewolson/testflight v1.0.0 // indirect
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/go-acme/lego v2.6.0+incompatible
	github.com/go-ini/ini v1.27.0 // indirect
	github.com/go-martini/martini v0.0.0-20170121215854-22fa46961aab // indirect
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/jarcoal/httpmock v1.0.4
	github.com/jinzhu/gorm v0.0.0-20160404144928-5174cc5c242a
	github.com/jinzhu/inflection v0.0.0-20170102125226-1c35d901db3d // indirect
	github.com/jinzhu/now v1.0.1 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20160803190731-bd40a432e4c7 // indirect
	github.com/kelseyhightower/envconfig v1.2.0
	github.com/lib/pq v0.0.0-20170324204654-2704adc878c2
	github.com/martini-contrib/render v0.0.0-20150707142108-ec18f8345a11 // indirect
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/miekg/dns v0.0.0-20170412184748-6ebcb714d369 // indirect
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pivotal-cf/brokerapi v0.0.0-20170307172350-8f25c50c2794
	github.com/robfig/cron v1.0.0
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	github.com/stretchr/testify v1.3.0
	github.com/xenolf/lego v2.6.0+incompatible
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	gopkg.in/square/go-jose.v1 v1.1.0 // indirect
	gopkg.in/square/go-jose.v2 v2.3.1 // indirect
)

replace github.com/xenolf/lego => github.com/jmcarp/lego v0.3.2-0.20170424160445-b4deb96f1082
