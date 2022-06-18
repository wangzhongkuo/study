package study

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
	"testing"
)

func TestNats(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// Create JetStream Context
	js, _ := nc.JetStream()

	js.Subscribe("OMNI.ORDERS.paid", func(m *nats.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
	}, nats.Durable("notify_task"))
}

var defaultConfigYamlString = []byte(`
kebab-case: chain case
`)

type Config struct {
	KebabCase string `mapstructure:"kebab-case"`
}

func Test_yaml(t *testing.T) {
	v := viper.New()
	// load default config
	v.SetConfigType("yaml")
	v.ReadConfig(bytes.NewBuffer(defaultConfigYamlString))
	conf := &Config{}
	err := v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(conf.KebabCase)
	fmt.Println(v.Get("kebab-case"))
}

func TestFloat(t *testing.T) {
	a := 1.010
	b := 1.01
	fmt.Println(a == b)
	i := uint64(1)
	fmt.Println(strconv.Itoa(int(i)))
}

func TestJson(t *testing.T) {
	receipt := `{
	    "OrderId":"GPA.3398-3452-4539-40792",
	    "packageName":"com.jssygp.bulletangel",
	    "productId":"com.sungray.dhlove.pack.52",
	    "purchaseTime":1653176973043,
	    "purchaseState":0,
	    "purchaseToken":"oidjdippfebblonakpmjoocm.AO-J1OyQzWC25OoT0RTSm8hPz45NvQQTGCF-4OEo45mLAyeI9gYHvHzdMohwkm91o8e5aptNMnAIkTjl7Hdnlb6s6Bm2kRIYZlpRWI2qn6AuORqodprp4VM",
	    "obfuscatedAccountId":"20436283",
	    "obfuscatedProfileId":"1528161077022601216",
	    "quantity":1,
	    "autoRenewing":true,
	    "acknowledged":false
	}`
	res := map[string]any{}
	json.Unmarshal([]byte(receipt), &res)
	fmt.Println(res)
	fmt.Println(res["autoRenewing"] == true)
	var gr *googleReceipt
	err := json.Unmarshal([]byte(receipt), &gr)
	fmt.Println(err)
	fmt.Println(gr.packageName)
}

type googleReceipt struct {
	OrderId             string `json:"orderId,omitempty"`
	packageName         string `json:"packageName,omitempty"`
	productId           string `json:"productId,omitempty"`
	purchaseTime        int64  `json:"purchaseTime,omitempty"`
	purchaseState       int64  `json:"purchaseState,omitempty"`
	purchaseToken       string `json:"purchaseToken,omitempty"`
	obfuscatedAccountId string `json:"obfuscatedAccountId,omitempty"`
	obfuscatedProfileId uint64 `json:"obfuscatedProfileId,omitempty"`
	quantity            int64  `json:"quantity,omitempty"`
	autoRenewing        bool   `json:"autoRenewing,omitempty"`
	acknowledged        bool   `json:"acknowledged,omitempty"`
}

type notificationType int64

const (
	SUBSCRIPTION_RECOVERED = notificationType(iota + 1)
	SUBSCRIPTION_RENEWED
	SUBSCRIPTION_CANCELED
	SUBSCRIPTION_PURCHASED
	SUBSCRIPTION_ON_HOLD
	SUBSCRIPTION_IN_GRACE_PERIOD
	SUBSCRIPTION_RESTARTED
	SUBSCRIPTION_PRICE_CHANGE_CONFIRMED
	SUBSCRIPTION_DEFERRED
	SUBSCRIPTION_PAUSED
	SUBSCRIPTION_PAUSE_SCHEDULE_CHANGED
	SUBSCRIPTION_REVOKED
	SUBSCRIPTION_EXPIRED
)

func TestIota(t *testing.T) {
	fmt.Println(reflect.TypeOf(SUBSCRIPTION_RENEWED))
}

func TestName(t *testing.T) {

}

type PostOpt interface {
	postValidate() error
}

type PostOptFn func() error

func (opt PostOptFn) postValidate() error {
	return opt()
}

func GoogleAck(a string) PostOpt {
	return PostOptFn(func() error {
		fmt.Println(a)
		return nil
	})
}
