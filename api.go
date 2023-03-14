package sys

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/jwt"
	"github.com/nats-io/nats.go"
)

const (
	DefaultRequestTimeout = 10 * time.Second
)

const (
	srvVarzSubj = "$SYS.REQ.SERVER.%s.VARZ"
)

// System can be used to request monitoring data from the server
type System struct {
	nc *nats.Conn
}

type VarzResp struct {
	Server ServerInfo `json:"server"`
	Varz   Varz       `json:"data"`
}

// ServerInfo identifies remote servers.
type ServerInfo struct {
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	ID        string    `json:"id"`
	Cluster   string    `json:"cluster,omitempty"`
	Domain    string    `json:"domain,omitempty"`
	Version   string    `json:"ver"`
	Tags      []string  `json:"tags,omitempty"`
	Seq       uint64    `json:"seq"`
	JetStream bool      `json:"jetstream"`
	Time      time.Time `json:"time"`
}

// VarzResp is a server response from VARZ endpoint, containing general information about the server.
type Varz struct {
	ID                    string                `json:"server_id"`
	Name                  string                `json:"server_name"`
	Version               string                `json:"version"`
	Proto                 int                   `json:"proto"`
	GitCommit             string                `json:"git_commit,omitempty"`
	GoVersion             string                `json:"go"`
	Host                  string                `json:"host"`
	Port                  int                   `json:"port"`
	AuthRequired          bool                  `json:"auth_required,omitempty"`
	TLSRequired           bool                  `json:"tls_required,omitempty"`
	TLSVerify             bool                  `json:"tls_verify,omitempty"`
	IP                    string                `json:"ip,omitempty"`
	ClientConnectURLs     []string              `json:"connect_urls,omitempty"`
	WSConnectURLs         []string              `json:"ws_connect_urls,omitempty"`
	MaxConn               int                   `json:"max_connections"`
	MaxSubs               int                   `json:"max_subscriptions,omitempty"`
	PingInterval          time.Duration         `json:"ping_interval"`
	MaxPingsOut           int                   `json:"ping_max"`
	HTTPHost              string                `json:"http_host"`
	HTTPPort              int                   `json:"http_port"`
	HTTPBasePath          string                `json:"http_base_path"`
	HTTPSPort             int                   `json:"https_port"`
	AuthTimeout           float64               `json:"auth_timeout"`
	MaxControlLine        int32                 `json:"max_control_line"`
	MaxPayload            int                   `json:"max_payload"`
	MaxPending            int64                 `json:"max_pending"`
	Cluster               ClusterOptsVarz       `json:"cluster,omitempty"`
	Gateway               GatewayOptsVarz       `json:"gateway,omitempty"`
	LeafNode              LeafNodeOptsVarz      `json:"leaf,omitempty"`
	MQTT                  MQTTOptsVarz          `json:"mqtt,omitempty"`
	Websocket             WebsocketOptsVarz     `json:"websocket,omitempty"`
	JetStream             JetStreamVarz         `json:"jetstream,omitempty"`
	TLSTimeout            float64               `json:"tls_timeout"`
	WriteDeadline         time.Duration         `json:"write_deadline"`
	Start                 time.Time             `json:"start"`
	Now                   time.Time             `json:"now"`
	Uptime                string                `json:"uptime"`
	Mem                   int64                 `json:"mem"`
	Cores                 int                   `json:"cores"`
	MaxProcs              int                   `json:"gomaxprocs"`
	CPU                   float64               `json:"cpu"`
	Connections           int                   `json:"connections"`
	TotalConnections      uint64                `json:"total_connections"`
	Routes                int                   `json:"routes"`
	Remotes               int                   `json:"remotes"`
	Leafs                 int                   `json:"leafnodes"`
	InMsgs                int64                 `json:"in_msgs"`
	OutMsgs               int64                 `json:"out_msgs"`
	InBytes               int64                 `json:"in_bytes"`
	OutBytes              int64                 `json:"out_bytes"`
	SlowConsumers         int64                 `json:"slow_consumers"`
	Subscriptions         uint32                `json:"subscriptions"`
	HTTPReqStats          map[string]uint64     `json:"http_req_stats"`
	ConfigLoadTime        time.Time             `json:"config_load_time"`
	Tags                  jwt.TagList           `json:"tags,omitempty"`
	TrustedOperatorsJwt   []string              `json:"trusted_operators_jwt,omitempty"`
	TrustedOperatorsClaim []*jwt.OperatorClaims `json:"trusted_operators_claim,omitempty"`
	SystemAccount         string                `json:"system_account,omitempty"`
	PinnedAccountFail     uint64                `json:"pinned_account_fails,omitempty"`
}

// ClusterOptsVarz contains monitoring cluster information
type ClusterOptsVarz struct {
	Name        string   `json:"name,omitempty"`
	Host        string   `json:"addr,omitempty"`
	Port        int      `json:"cluster_port,omitempty"`
	AuthTimeout float64  `json:"auth_timeout,omitempty"`
	URLs        []string `json:"urls,omitempty"`
	TLSTimeout  float64  `json:"tls_timeout,omitempty"`
	TLSRequired bool     `json:"tls_required,omitempty"`
	TLSVerify   bool     `json:"tls_verify,omitempty"`
}

// GatewayOptsVarz contains monitoring gateway information
type GatewayOptsVarz struct {
	Name           string                  `json:"name,omitempty"`
	Host           string                  `json:"host,omitempty"`
	Port           int                     `json:"port,omitempty"`
	AuthTimeout    float64                 `json:"auth_timeout,omitempty"`
	TLSTimeout     float64                 `json:"tls_timeout,omitempty"`
	TLSRequired    bool                    `json:"tls_required,omitempty"`
	TLSVerify      bool                    `json:"tls_verify,omitempty"`
	Advertise      string                  `json:"advertise,omitempty"`
	ConnectRetries int                     `json:"connect_retries,omitempty"`
	Gateways       []RemoteGatewayOptsVarz `json:"gateways,omitempty"`
	RejectUnknown  bool                    `json:"reject_unknown,omitempty"` // config got renamed to reject_unknown_cluster
}

// RemoteGatewayOptsVarz contains monitoring remote gateway information
type RemoteGatewayOptsVarz struct {
	Name       string   `json:"name"`
	TLSTimeout float64  `json:"tls_timeout,omitempty"`
	URLs       []string `json:"urls,omitempty"`
}

// LeafNodeOptsVarz contains monitoring leaf node information
type LeafNodeOptsVarz struct {
	Host        string               `json:"host,omitempty"`
	Port        int                  `json:"port,omitempty"`
	AuthTimeout float64              `json:"auth_timeout,omitempty"`
	TLSTimeout  float64              `json:"tls_timeout,omitempty"`
	TLSRequired bool                 `json:"tls_required,omitempty"`
	TLSVerify   bool                 `json:"tls_verify,omitempty"`
	Remotes     []RemoteLeafOptsVarz `json:"remotes,omitempty"`
}

// RemoteLeafOptsVarz contains monitoring remote leaf node information
type RemoteLeafOptsVarz struct {
	LocalAccount string     `json:"local_account,omitempty"`
	TLSTimeout   float64    `json:"tls_timeout,omitempty"`
	URLs         []string   `json:"urls,omitempty"`
	Deny         *DenyRules `json:"deny,omitempty"`
}

// DenyRules Contains lists of subjects not allowed to be imported/exported
type DenyRules struct {
	Exports []string `json:"exports,omitempty"`
	Imports []string `json:"imports,omitempty"`
}

// MQTTOptsVarz contains monitoring MQTT information
type MQTTOptsVarz struct {
	Host           string        `json:"host,omitempty"`
	Port           int           `json:"port,omitempty"`
	NoAuthUser     string        `json:"no_auth_user,omitempty"`
	AuthTimeout    float64       `json:"auth_timeout,omitempty"`
	TLSMap         bool          `json:"tls_map,omitempty"`
	TLSTimeout     float64       `json:"tls_timeout,omitempty"`
	TLSPinnedCerts []string      `json:"tls_pinned_certs,omitempty"`
	JsDomain       string        `json:"js_domain,omitempty"`
	AckWait        time.Duration `json:"ack_wait,omitempty"`
	MaxAckPending  uint16        `json:"max_ack_pending,omitempty"`
}

// WebsocketOptsVarz contains monitoring websocket information
type WebsocketOptsVarz struct {
	Host             string        `json:"host,omitempty"`
	Port             int           `json:"port,omitempty"`
	Advertise        string        `json:"advertise,omitempty"`
	NoAuthUser       string        `json:"no_auth_user,omitempty"`
	JWTCookie        string        `json:"jwt_cookie,omitempty"`
	HandshakeTimeout time.Duration `json:"handshake_timeout,omitempty"`
	AuthTimeout      float64       `json:"auth_timeout,omitempty"`
	NoTLS            bool          `json:"no_tls,omitempty"`
	TLSMap           bool          `json:"tls_map,omitempty"`
	TLSPinnedCerts   []string      `json:"tls_pinned_certs,omitempty"`
	SameOrigin       bool          `json:"same_origin,omitempty"`
	AllowedOrigins   []string      `json:"allowed_origins,omitempty"`
	Compression      bool          `json:"compression,omitempty"`
}

// JetStreamVarz contains basic runtime information about jetstream
type JetStreamVarz struct {
	Config *JetStreamConfig `json:"config,omitempty"`
	Stats  *JetStreamStats  `json:"stats,omitempty"`
	Meta   *MetaClusterInfo `json:"meta,omitempty"`
}

// Statistics about JetStream for this server.
type JetStreamStats struct {
	Memory         uint64            `json:"memory"`
	Store          uint64            `json:"storage"`
	ReservedMemory uint64            `json:"reserved_memory"`
	ReservedStore  uint64            `json:"reserved_storage"`
	Accounts       int               `json:"accounts"`
	HAAssets       int               `json:"ha_assets"`
	API            JetStreamAPIStats `json:"api"`
}

// JetStreamConfig determines this server's configuration.
// MaxMemory and MaxStore are in bytes.
type JetStreamConfig struct {
	MaxMemory  int64  `json:"max_memory"`
	MaxStore   int64  `json:"max_storage"`
	StoreDir   string `json:"store_dir,omitempty"`
	Domain     string `json:"domain,omitempty"`
	CompressOK bool   `json:"compress_ok,omitempty"`
	UniqueTag  string `json:"unique_tag,omitempty"`
}

type JetStreamAPIStats struct {
	Total    uint64 `json:"total"`
	Errors   uint64 `json:"errors"`
	Inflight uint64 `json:"inflight,omitempty"`
}

// MetaClusterInfo shows information about the meta group.
type MetaClusterInfo struct {
	Name     string      `json:"name,omitempty"`
	Leader   string      `json:"leader,omitempty"`
	Peer     string      `json:"peer,omitempty"`
	Replicas []*PeerInfo `json:"replicas,omitempty"`
	Size     int         `json:"cluster_size"`
}

// PeerInfo shows information about all the peers in the cluster that
// are supporting the stream or consumer.
type PeerInfo struct {
	Name    string        `json:"name"`
	Current bool          `json:"current"`
	Offline bool          `json:"offline,omitempty"`
	Active  time.Duration `json:"active"`
	Lag     uint64        `json:"lag,omitempty"`
	Peer    string        `json:"peer"`
}

// In the context of system events, VarzEventOptions are options passed to Varz
type VarzEventOptions struct {
	EventFilterOptions
}

// Common filter options for system requests STATSZ VARZ SUBSZ CONNZ ROUTEZ GATEWAYZ LEAFZ
type EventFilterOptions struct {
	Name    string   `json:"server_name,omitempty"` // filter by server name
	Cluster string   `json:"cluster,omitempty"`     // filter by cluster name
	Host    string   `json:"host,omitempty"`        // filter by host name
	Tags    []string `json:"tags,omitempty"`        // filter by tags (must match all tags)
	Domain  string   `json:"domain,omitempty"`      // filter by JS domain
}

func NewSysClient(nc *nats.Conn) System {
	return System{
		nc: nc,
	}
}

// Varz returns general server information
func (s *System) Varz(id string, opts VarzEventOptions) (*VarzResp, error) {
	conn := s.nc
	subj := fmt.Sprintf(srvVarzSubj, id)
	payload, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	resp, err := conn.Request(subj, payload, DefaultRequestTimeout)
	if err != nil {
		return nil, err
	}

	var varzResp VarzResp
	if err := json.Unmarshal(resp.Data, &varzResp); err != nil {
		return nil, err
	}
	fmt.Println(string(resp.Data))
	fmt.Println()

	return &varzResp, nil
}

func (s *System) VarzPing(opts VarzEventOptions) ([]VarzResp, error) {
	subj := fmt.Sprintf(srvVarzSubj, "PING")
	payload, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	resp, err := s.RequestMany(subj, payload)
	if err != nil {
		return nil, err
	}
	srvVarz := make([]VarzResp, 0, len(resp))
	for _, msg := range resp {
		var varzResp VarzResp
		if err := json.Unmarshal(msg.Data, &varzResp); err != nil {
			return nil, err
		}
		srvVarz = append(srvVarz, varzResp)
	}
	return srvVarz, nil
}

type requestManyOpts struct {
	maxWait     time.Duration
	maxInterval time.Duration
	count       int
}

type RequestManyOpt func(*requestManyOpts) error

func WithRequestManyMaxWait(maxWait time.Duration) RequestManyOpt {
	return func(opts *requestManyOpts) error {
		opts.maxWait = maxWait
		return nil
	}
}

func WithRequestManyMaxInterval(interval time.Duration) RequestManyOpt {
	return func(opts *requestManyOpts) error {
		opts.maxInterval = interval
		return nil
	}
}

func WithRequestManyCount(count int) RequestManyOpt {
	return func(opts *requestManyOpts) error {
		opts.count = count
		return nil
	}
}

func (s *System) RequestMany(subject string, data []byte, opts ...RequestManyOpt) ([]*nats.Msg, error) {
	conn := s.nc
	reqOpts := &requestManyOpts{
		maxWait:     DefaultRequestTimeout,
		maxInterval: 300 * time.Millisecond,
		count:       -1,
	}

	for _, opt := range opts {
		if err := opt(reqOpts); err != nil {
			return nil, err
		}
	}

	inbox := nats.NewInbox()
	res := make([]*nats.Msg, 0)
	msgsChan := make(chan *nats.Msg, 100)

	intervalTimer := time.NewTimer(reqOpts.maxInterval)
	sub, err := conn.Subscribe(inbox, func(msg *nats.Msg) {
		intervalTimer.Reset(reqOpts.maxInterval)
		msgsChan <- msg
	})
	defer sub.Unsubscribe()

	if err := conn.PublishRequest(subject, inbox, data); err != nil {
		return nil, err
	}

	for {
		select {
		case msg := <-msgsChan:
			if msg.Header.Get("Status") == "503" {
				return nil, fmt.Errorf("server request on subject %q failed: %s", subject, err)
			}
			res = append(res, msg)
			if reqOpts.count != -1 && len(res) == reqOpts.count {
				return res, nil
			}
		case <-intervalTimer.C:
			return res, nil
		case <-time.After(reqOpts.maxWait):
			return res, nil
		}
	}
}
