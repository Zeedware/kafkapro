package main

type KafkaConfiguration struct {
	key     string
	topic   string
	groupId string
	ip      string
	port    string
}

func (k *KafkaConfiguration) Ip() string {
	return k.ip
}

func (k *KafkaConfiguration) SetIp(ip string) {
	k.ip = ip
}

func (k *KafkaConfiguration) GroupId() string {
	return k.groupId
}

func (k *KafkaConfiguration) SetGroupId(groupId string) {
	k.groupId = groupId
}

func (k *KafkaConfiguration) Topic() string {
	return k.topic
}

func (k *KafkaConfiguration) SetTopic(topic string) {
	k.topic = topic
}

func (k *KafkaConfiguration) Port() string {
	return k.port
}

func (k *KafkaConfiguration) SetPort(port string) {
	k.port = port
}

func (k *KafkaConfiguration) Key() string {
	return k.key
}

func (k *KafkaConfiguration) SetKey(key string) {
	k.key = key
}

func (k *KafkaConfiguration) Host() string {
	return k.Ip() + ":" + k.Port()
}

func NewKafkaConfiguration(key string, topic string, groupId string, ip string, port string) *KafkaConfiguration {
	return &KafkaConfiguration{key: key, topic: topic, groupId: groupId, ip: ip, port: port}
}
