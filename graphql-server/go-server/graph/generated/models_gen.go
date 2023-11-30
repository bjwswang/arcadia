// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"time"
)

type PageNode interface {
	IsPageNode()
}

type AddDataProcessInput struct {
	Name                  string                   `json:"name"`
	FileType              string                   `json:"file_type"`
	PreDataSetName        string                   `json:"pre_data_set_name"`
	PreDataSetVersion     string                   `json:"pre_data_set_version"`
	FileNames             []*FileItem              `json:"file_names,omitempty"`
	PostDataSetName       string                   `json:"post_data_set_name"`
	PostDataSetVersion    string                   `json:"post_data_set_version"`
	DataProcessConfigInfo []*DataProcessConfigItem `json:"data_process_config_info,omitempty"`
	BucketName            string                   `json:"bucket_name"`
	VersionDataSetName    string                   `json:"version_data_set_name"`
}

type AllDataProcessListByCountInput struct {
	Keyword string `json:"keyword"`
}

type AllDataProcessListByPageInput struct {
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Keyword   string `json:"keyword"`
}

type CountDataProcessItem struct {
	Status  int    `json:"status"`
	Data    int    `json:"data"`
	Message string `json:"message"`
}

// 数据集创建的输入
type CreateDatasetInput struct {
	// 数据集的名字
	// 规则: k8s的名称规则
	// 规则: 非空
	Name string `json:"name"`
	// 数据集的命名空间
	// 规则: 非空
	Namespace string `json:"namespace"`
	// 一些标签选择信息，可以不添加
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 一些备注用的注视信息，或者记录一个简单的配置
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty"`
	// 描述信息，可以不写
	Description *string `json:"description,omitempty"`
	// 数据集里面的数据的类型，文本，视频，图片
	ContentType string `json:"contentType"`
	// 应用场景，可以为空
	// 规则: enum{ xx xx } (非固定字段，以产品为准)
	Filed *string `json:"filed,omitempty"`
}

// 新增数据源时输入条件
type CreateDatasourceInput struct {
	// 名字
	// 规则: k8s的名称规则
	// 规则: 非空
	Name string `json:"name"`
	// 数据源的命名空间
	// 规则: 非空
	Namespace string `json:"namespace"`
	// 数据源资源标签
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 数据源资源注释
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 数据源资源展示名称作为显示，并提供编辑
	DisplayName *string `json:"displayName,omitempty"`
	// 数据源资源描述
	Description *string `json:"description,omitempty"`
	// 提供对象存储时输入条件
	Endpointinput *EndpointInput `json:"endpointinput,omitempty"`
	Ossinput      *OssInput      `json:"ossinput,omitempty"`
}

type CreateEmbedderInput struct {
	// 模型服务资源名称（不可同名）
	Name string `json:"name"`
	// 模型服务创建命名空间
	Namespace string `json:"namespace"`
	// 模型服务资源标签
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 模型服务资源注释
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 模型服务资源展示名称作为显示，并提供编辑
	DisplayName *string `json:"displayName,omitempty"`
	// 模型服务资源描述
	Description   *string        `json:"description,omitempty"`
	Endpointinput *EndpointInput `json:"endpointinput,omitempty"`
	// 模型服务类型
	ServiceType *string `json:"serviceType,omitempty"`
}

// 创建知识库的输入
type CreateKnowledgeBaseInput struct {
	// 知识库资源名称（不可同名）
	Name string `json:"name"`
	// 知识库创建命名空间
	Namespace string `json:"namespace"`
	// 知识库资源标签
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 知识库资源注释
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 知识库资源展示名称作为显示，并提供编辑
	DisplayName *string `json:"displayName,omitempty"`
	// 知识库资源描述
	Description *string `json:"description,omitempty"`
	// embedder指当前知识库使用的embedding向量化模型
	Embedder string `json:"embedder"`
	// "向量数据库(目前不需要填写，直接使用系统默认的向量数据库)
	VectorStore *TypedObjectReferenceInput `json:"vectorStore,omitempty"`
	// 知识库文件
	FileGroups []*Filegroupinput `json:"fileGroups,omitempty"`
}

type CreateModelInput struct {
	// 模型资源名称（不可同名）
	Name string `json:"name"`
	// 模型创建命名空间
	Namespace string `json:"namespace"`
	// 模型资源展示名称作为显示，并提供编辑
	DisplayName *string `json:"displayName,omitempty"`
	// 模型资源描述
	Description *string `json:"description,omitempty"`
	// 模型类型
	Modeltypes string `json:"modeltypes"`
}

type CreateVersionedDatasetInput struct {
	// 数据集的CR名字，要满足k8s的名称规则
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	// dataset的名字，需要根据这个名字，
	//     判断是否最新版本不包含任何文件(产品要求，有一个不包含任何文件的版本，不允许创建新的版本)
	DatasetName string `json:"datasetName"`
	// 一些标签选择信息，可以不添加
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 一些备注用的注视信息，或者记录一个简单的配置
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 展示名称，用于展示在界面上的，必须填写
	DisplayName *string `json:"displayName,omitempty"`
	// 描述信息，可以不写
	Description *string `json:"description,omitempty"`
	// 数据集里面的数据的类型，文本，视频，图片
	Version string `json:"version"`
	// 是否发布，0是未发布，1是已经发布，创建一个版本的时候默认传递0就可以
	Released int `json:"released"`
	// 从数据源要上传的文件，目前以及不用了
	FileGrups []*FileGroup `json:"fileGrups,omitempty"`
	// 界面上创建新版本选择从某个版本集成的时候，填写version字段
	InheritedFrom *string `json:"inheritedFrom,omitempty"`
}

type DataProcessConfig struct {
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Status      string                       `json:"status"`
	Children    []*DataProcessConfigChildren `json:"children,omitempty"`
}

type DataProcessConfigChildren struct {
	Name        *string                     `json:"name,omitempty"`
	Enable      *string                     `json:"enable,omitempty"`
	ZhName      *string                     `json:"zh_name,omitempty"`
	Description *string                     `json:"description,omitempty"`
	Preview     []*DataProcessConfigpreView `json:"preview,omitempty"`
}

type DataProcessConfigItem struct {
	Type string `json:"type"`
}

type DataProcessConfigpreView struct {
	FileName *string                            `json:"file_name,omitempty"`
	Content  []*DataProcessConfigpreViewContent `json:"content,omitempty"`
}

type DataProcessConfigpreViewContent struct {
	Pre  *string `json:"pre,omitempty"`
	Post *string `json:"post,omitempty"`
}

type DataProcessDetails struct {
	Status  int                    `json:"status"`
	Data    DataProcessDetailsItem `json:"data"`
	Message string                 `json:"message"`
}

type DataProcessDetailsInput struct {
	ID string `json:"id"`
}

type DataProcessDetailsItem struct {
	ID                 string               `json:"id"`
	Status             string               `json:"status"`
	FileType           string               `json:"file_type"`
	PreDatasetName     string               `json:"pre_dataset_name"`
	PreDatasetVersion  string               `json:"pre_dataset_version"`
	PostDatasetName    string               `json:"post_dataset_name"`
	PostDatasetVersion string               `json:"post_dataset_version"`
	FileNum            int                  `json:"file_num"`
	StartTime          string               `json:"start_time"`
	EndTime            string               `json:"end_time"`
	Config             []*DataProcessConfig `json:"config,omitempty"`
}

type DataProcessItem struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Status             string  `json:"status"`
	PreDataSetName     string  `json:"pre_data_set_name"`
	PreDataSetVersion  string  `json:"pre_data_set_version"`
	PostDataSetName    string  `json:"post_data_set_name"`
	PostDataSetVersion *string `json:"post_data_set_version,omitempty"`
	StartDatetime      string  `json:"start_datetime"`
}

type DataProcessMutation struct {
	CreateDataProcessTask *DataProcessResponse `json:"createDataProcessTask,omitempty"`
	DeleteDataProcessTask *DataProcessResponse `json:"deleteDataProcessTask,omitempty"`
}

type DataProcessQuery struct {
	AllDataProcessListByPage  *PaginatedDataProcessItem `json:"allDataProcessListByPage,omitempty"`
	AllDataProcessListByCount *CountDataProcessItem     `json:"allDataProcessListByCount,omitempty"`
	DataProcessSupportType    *DataProcessSupportType   `json:"dataProcessSupportType,omitempty"`
	DataProcessDetails        *DataProcessDetails       `json:"dataProcessDetails,omitempty"`
}

type DataProcessResponse struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

type DataProcessSupportType struct {
	Status  int                           `json:"status"`
	Data    []*DataProcessSupportTypeItem `json:"data,omitempty"`
	Message string                        `json:"message"`
}

type DataProcessSupportTypeChildren struct {
	Name        string `json:"name"`
	ZhName      string `json:"zh_name"`
	Enable      string `json:"enable"`
	Description string `json:"description"`
}

type DataProcessSupportTypeItem struct {
	Name        string                            `json:"name"`
	Description string                            `json:"description"`
	Children    []*DataProcessSupportTypeChildren `json:"children,omitempty"`
}

// Dataset
// 数据集代表用户纳管的一组相似属性的文件，采用相同的方式进行数据处理并用于后续的
// 1. 模型训练
// 2. 知识库
//
// 支持多种类型数据:
// - 文本
// - 图片
// - 视频
//
// 单个数据集仅允许包含同一类型文件，不同类型文件将被忽略
// 数据集允许有多个版本，数据处理针对单个版本进行
// 数据集某个版本完成数据处理后，数据处理服务需要将处理后的存储回 版本数据集
type Dataset struct {
	// 名称
	// 规则: 遵循k8s命名
	Name string `json:"name"`
	// 所在的namespace(文件上传时作为bucket)
	// 规则: 获取当前项目对应的命名空间
	// 规则: 非空
	Namespace string `json:"namespace"`
	// 一些用于标记，选择的的标签
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 添加一些辅助性记录信息
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 创建者，为当前用户的用户名
	// 规则: webhook启用后自动添加，默认为空
	Creator *string `json:"creator,omitempty"`
	// 展示名
	DisplayName *string `json:"displayName,omitempty"`
	// 描述信息
	Description *string `json:"description,omitempty"`
	// 创建时间
	CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
	// 更新时间, 这里更新指文件同步，或者数据处理完成后，做的更新操作的时间
	UpdateTimestamp *time.Time `json:"updateTimestamp,omitempty"`
	// 数据集类型，文本，图片，视频
	// 规则: enum{ text image video}
	// 规则: 非空
	ContentType string `json:"contentType"`
	// 应用场景
	// 规则: enum{ xx xx } (非固定字段，以产品为准)
	Field *string `json:"field,omitempty"`
	// 数据集下面的版本列表。
	// 规则: 支持对名字，类型的完全匹配过滤。
	// 规则: 支持通过标签(somelabel=abc)，字段(metadata.name=abc)进行过滤
	Versions PaginatedResult `json:"versions"`
}

func (Dataset) IsPageNode() {}

// 数据集更新
type DatasetMutation struct {
	// 创建数据集
	CreateDataset Dataset `json:"createDataset"`
	// 更新数据集
	UpdateDataset Dataset `json:"updateDataset"`
	// 删除数据集
	// 规则: 支持删除一个名称列表中包含的所有数据集
	// 规则: 支持通过标签选择器，将满足标签的dataset全部删除
	// 规则: 如果提供了这两个参数，以名字列表为主。
	DeleteDatasets *string `json:"deleteDatasets,omitempty"`
}

// 数据集查询
type DatasetQuery struct {
	// 根据名字获取某个具体的数据集
	GetDataset Dataset `json:"getDataset"`
	// 获取数据集列表
	// 规则: 支持通过标签和字段进行选择。如下:
	// labelSelector: aa=bbb
	// fieldSelector= metadata.name=somename
	ListDatasets PaginatedResult `json:"listDatasets"`
}

// 数据源: 定义了对某一个具备数据存储能力服务的访问信息，供后续向该数据源获取数据使用
type Datasource struct {
	// 名称
	// 规则: 遵循k8s命名
	// 规则: 非空
	Name string `json:"name"`
	// 命名空间
	// 规则: 非空
	Namespace   string                 `json:"namespace"`
	Labels      map[string]interface{} `json:"labels,omitempty"`
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 创建者，为当前用户的用户名
	// 规则: webhook启用后自动添加，默认为空
	Creator *string `json:"creator,omitempty"`
	// 展示名
	DisplayName *string `json:"displayName,omitempty"`
	// 描述信息
	Description *string `json:"description,omitempty"`
	// 终端访问信息
	Endpoint *Endpoint `json:"endpoint,omitempty"`
	// 对象存储访问信息
	// 规则: 非空代表当前数据源为对象存储数据源
	Oss *Oss `json:"oss,omitempty"`
	// 数据源连接状态
	Status *string `json:"status,omitempty"`
	// 文件数量
	FileCount *int `json:"fileCount,omitempty"`
	// 创建时间
	CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
	// 更新时间, 这里更新指文件同步，或者数据处理完成后，做的更新操作的时间
	UpdateTimestamp *time.Time `json:"updateTimestamp,omitempty"`
}

func (Datasource) IsPageNode() {}

type DatasourceMutation struct {
	CreateDatasource Datasource `json:"createDatasource"`
	UpdateDatasource Datasource `json:"updateDatasource"`
	DeleteDatasource *string    `json:"deleteDatasource,omitempty"`
}

type DatasourceQuery struct {
	GetDatasource   Datasource      `json:"getDatasource"`
	ListDatasources PaginatedResult `json:"listDatasources"`
}

type DeleteDataProcessInput struct {
	ID string `json:"id"`
}

// 数据集删除的输入
type DeleteDatasetInput struct {
	// name, namespace用来确定资源
	Name      *string `json:"name,omitempty"`
	Namespace string  `json:"namespace"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
}

// 删除数据源的输入
type DeleteDatasourceInput struct {
	// name, namespace用来确定资源
	Name      *string `json:"name,omitempty"`
	Namespace string  `json:"namespace"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
}

type DeleteEmbedderInput struct {
	Name      *string `json:"name,omitempty"`
	Namespace string  `json:"namespace"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
}

// 知识库删除的输入
type DeleteKnowledgeBaseInput struct {
	Name      *string `json:"name,omitempty"`
	Namespace string  `json:"namespace"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
}

type DeleteModelInput struct {
	Name      *string `json:"name,omitempty"`
	Namespace string  `json:"namespace"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
}

type DeleteVersionedDatasetInput struct {
	Name          *string `json:"name,omitempty"`
	Namespace     string  `json:"namespace"`
	LabelSelector *string `json:"labelSelector,omitempty"`
	FieldSelector *string `json:"fieldSelector,omitempty"`
}

type Embedder struct {
	Name            string                 `json:"name"`
	Namespace       string                 `json:"namespace"`
	Labels          map[string]interface{} `json:"labels,omitempty"`
	Annotations     map[string]interface{} `json:"annotations,omitempty"`
	Creator         *string                `json:"creator,omitempty"`
	DisplayName     *string                `json:"displayName,omitempty"`
	Description     *string                `json:"description,omitempty"`
	Endpoint        *Endpoint              `json:"endpoint,omitempty"`
	ServiceType     *string                `json:"serviceType,omitempty"`
	UpdateTimestamp *time.Time             `json:"updateTimestamp,omitempty"`
}

func (Embedder) IsPageNode() {}

type EmbedderMutation struct {
	CreateEmbedder Embedder `json:"createEmbedder"`
	UpdateEmbedder Embedder `json:"updateEmbedder"`
	DeleteEmbedder *string  `json:"deleteEmbedder,omitempty"`
}

type EmbedderQuery struct {
	GetEmbedder   Embedder        `json:"getEmbedder"`
	ListEmbedders PaginatedResult `json:"listEmbedders"`
}

// 终端的访问信息
type Endpoint struct {
	// url地址
	URL *string `json:"url,omitempty"`
	// 终端访问的密钥信息，保存在k8s secret中
	AuthSecret *TypedObjectReference `json:"authSecret,omitempty"`
	// 是否通过非安全方式访问，默认为false，即安全模式访问
	Insecure *bool `json:"insecure,omitempty"`
}

// 对象存储终端输入
type EndpointInput struct {
	URL *string `json:"url,omitempty"`
	// secret验证密码
	AuthSecret *TypedObjectReferenceInput `json:"authSecret,omitempty"`
	// 默认true
	Insecure *bool `json:"insecure,omitempty"`
}

// File
// 展示某个版本的所有文件。
type F struct {
	// 文件在数据源中的路径，a/b/c.txt或者d.txt
	Path string `json:"path"`
	// 文件类型
	FileType string `json:"fileType"`
	// 数据量
	Count *int `json:"count,omitempty"`
	// 文件成功导入时间，如果没有导入成功，这个字段为空
	Time *time.Time `json:"time,omitempty"`
}

func (F) IsPageNode() {}

// 根据条件顾虑版本内的文件，只支持关键词搜索
type FileFilter struct {
	// 根据关键词搜索文件，strings.Container(fileName, keyword)
	Keyword *string `json:"keyword,omitempty"`
	// 页
	Page *int `json:"page,omitempty"`
	// 页内容数量
	PageSize *int `json:"pageSize,omitempty"`
	// 根据文件名字或者更新时间排序, file, time
	SortBy *string `json:"sortBy,omitempty"`
}

type FileGroup struct {
	// 数据源的基础信息
	Source TypedObjectReferenceInput `json:"source"`
	// 用到的文件路径，注意⚠️ 一定不要加bucket的名字
	Paths []string `json:"paths,omitempty"`
}

type FileItem struct {
	Name string `json:"name"`
}

// 知识库
type KnowledgeBase struct {
	// 知识库id,为CR资源中的metadata.uid
	ID *string `json:"id,omitempty"`
	// 名称
	// 规则: 遵循k8s命名
	Name string `json:"name"`
	// 所在的namespace(文件上传时作为bucket)
	// 规则: 获取当前项目对应的命名空间
	// 规则: 非空
	Namespace string `json:"namespace"`
	// 一些用于标记，选择的的标签
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 添加一些辅助性记录信息
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 创建者，为当前用户的用户名
	// 规则: webhook启用后自动添加，默认为空
	Creator *string `json:"creator,omitempty"`
	// 展示名
	DisplayName *string `json:"displayName,omitempty"`
	// 描述信息
	Description *string `json:"description,omitempty"`
	// 创建时间
	CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
	// 更新时间
	UpdateTimestamp *time.Time `json:"updateTimestamp,omitempty"`
	// embedder指当前知识库使用的embedding向量化模型，即 Kind 为 Embedder
	Embedder *TypedObjectReference `json:"embedder,omitempty"`
	// vectorStore指当前知识库使用的向量数据库服务，即 Kind 为 VectorStore
	VectorStore *TypedObjectReference `json:"vectorStore,omitempty"`
	// fileGroupDetails为知识库中所处理的文件组的详细内容和状态
	FileGroupDetails []*Filegroupdetail `json:"fileGroupDetails,omitempty"`
	// 知识库整体连接状态
	Status *string `json:"status,omitempty"`
}

func (KnowledgeBase) IsPageNode() {}

type KnowledgeBaseMutation struct {
	CreateKnowledgeBase KnowledgeBase `json:"createKnowledgeBase"`
	UpdateKnowledgeBase KnowledgeBase `json:"updateKnowledgeBase"`
	DeleteKnowledgeBase *string       `json:"deleteKnowledgeBase,omitempty"`
}

type KnowledgeBaseQuery struct {
	GetKnowledgeBase   KnowledgeBase   `json:"getKnowledgeBase"`
	ListKnowledgeBases PaginatedResult `json:"listKnowledgeBases"`
}

// 数据集分页列表查询的输入
type ListDatasetInput struct {
	// namespace用来确定资源
	// 规则: 必填
	Namespace string `json:"namespace"`
	// name用来唯一确定资源
	Name *string `json:"name,omitempty"`
	// 展示名
	DisplayName *string `json:"displayName,omitempty"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
	// 分页页码，
	// 规则: 从1开始，默认是1
	Page *int `json:"page,omitempty"`
	// 每页数量，
	// 规则: 默认10
	PageSize *int `json:"pageSize,omitempty"`
	// 关键词: 模糊匹配
	// 规则: namespace,name,displayName,contentType,annotations中如果包含该字段则返回
	Keyword *string `json:"keyword,omitempty"`
}

// 数据源分页列表查询的输入
type ListDatasourceInput struct {
	// namespace用来确定资源
	// 规则: 必填
	Namespace string `json:"namespace"`
	// name用来唯一确定资源
	Name *string `json:"name,omitempty"`
	// 展示名
	DisplayName *string `json:"displayName,omitempty"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
	// 分页页码，
	// 规则: 从1开始，默认是1
	Page *int `json:"page,omitempty"`
	// 每页数量，
	// 规则: 默认10
	PageSize *int `json:"pageSize,omitempty"`
	// 关键词: 模糊匹配
	// 规则: namespace,name,displayName,contentType,annotations中如果包含该字段则返回
	Keyword *string `json:"keyword,omitempty"`
}

type ListEmbedderInput struct {
	Name        *string `json:"name,omitempty"`
	Namespace   string  `json:"namespace"`
	DisplayName *string `json:"displayName,omitempty"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
	Page          *int    `json:"page,omitempty"`
	PageSize      *int    `json:"pageSize,omitempty"`
	Keyword       *string `json:"keyword,omitempty"`
}

// 知识库分页列表查询的输入
type ListKnowledgeBaseInput struct {
	Name        *string `json:"name,omitempty"`
	Namespace   string  `json:"namespace"`
	DisplayName *string `json:"displayName,omitempty"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
	// 分页页码，
	// 规则: 从1开始，默认是1
	Page *int `json:"page,omitempty"`
	// 每页数量，
	// 规则: 默认10
	PageSize *int `json:"pageSize,omitempty"`
	// 关键词: 模糊匹配
	// 规则: name,displayName中如果包含该字段则返回
	Keyword *string `json:"keyword,omitempty"`
}

type ListModelInput struct {
	Name        *string `json:"name,omitempty"`
	Namespace   string  `json:"namespace"`
	DisplayName *string `json:"displayName,omitempty"`
	// 标签选择器
	LabelSelector *string `json:"labelSelector,omitempty"`
	// 字段选择器
	FieldSelector *string `json:"fieldSelector,omitempty"`
	Page          *int    `json:"page,omitempty"`
	PageSize      *int    `json:"pageSize,omitempty"`
	Keyword       *string `json:"keyword,omitempty"`
}

type ListVersionedDatasetInput struct {
	Name          *string `json:"name,omitempty"`
	Namespace     *string `json:"namespace,omitempty"`
	DisplayName   *string `json:"displayName,omitempty"`
	LabelSelector *string `json:"labelSelector,omitempty"`
	FieldSelector *string `json:"fieldSelector,omitempty"`
	// 分页页码，从1开始，默认是1
	Page *int `json:"page,omitempty"`
	// 每页数量，默认10
	PageSize *int    `json:"pageSize,omitempty"`
	Keyword  *string `json:"keyword,omitempty"`
}

type Model struct {
	ID                *string                `json:"id,omitempty"`
	Name              string                 `json:"name"`
	Namespace         string                 `json:"namespace"`
	Labels            map[string]interface{} `json:"labels,omitempty"`
	Annotations       map[string]interface{} `json:"annotations,omitempty"`
	Creator           *string                `json:"creator,omitempty"`
	DisplayName       *string                `json:"displayName,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Modeltypes        string                 `json:"modeltypes"`
	Status            *string                `json:"status,omitempty"`
	CreationTimestamp *time.Time             `json:"creationTimestamp,omitempty"`
	UpdateTimestamp   *time.Time             `json:"updateTimestamp,omitempty"`
}

func (Model) IsPageNode() {}

type ModelMutation struct {
	CreateModel Model   `json:"createModel"`
	UpdateModel Model   `json:"updateModel"`
	DeleteModel *string `json:"deleteModel,omitempty"`
}

type ModelQuery struct {
	GetModel   Model           `json:"getModel"`
	ListModels PaginatedResult `json:"listModels"`
}

// 对象存储的使用信息
type Oss struct {
	// 所用的bucket名称
	Bucket *string `json:"bucket,omitempty"`
	// 所用的object路径(可为前缀)
	Object *string `json:"Object,omitempty"`
}

// 文件输入
type OssInput struct {
	Bucket *string `json:"bucket,omitempty"`
	Object *string `json:"Object,omitempty"`
}

type PaginatedDataProcessItem struct {
	Status  int                `json:"status"`
	Data    []*DataProcessItem `json:"data,omitempty"`
	Message string             `json:"message"`
}

type PaginatedResult struct {
	HasNextPage bool       `json:"hasNextPage"`
	Nodes       []PageNode `json:"nodes,omitempty"`
	Page        *int       `json:"page,omitempty"`
	PageSize    *int       `json:"pageSize,omitempty"`
	TotalCount  int        `json:"totalCount"`
}

type TypedObjectReference struct {
	APIGroup  *string `json:"apiGroup,omitempty"`
	Kind      string  `json:"kind"`
	Name      string  `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
}

type TypedObjectReferenceInput struct {
	APIGroup  *string `json:"apiGroup,omitempty"`
	Kind      string  `json:"kind"`
	Name      string  `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
}

// 数据集更新的输入
type UpdateDatasetInput struct {
	// name, namespace用来确定资源
	// 规则: 不允许修改的。将原数据传递回来即可。
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	// 更新的的标签信息，这里涉及到增加或者删除标签，
	// 规则: 不允许修改的。将原数据传递回来即可。
	// 如果标签有任何改动，传递完整的label。
	// 例如之前的标齐是: abc:def 新增一个标签aa:bb, 那么传递 abc:def, aa:bb
	Labels      map[string]interface{} `json:"labels,omitempty"`
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 如不更新，则为空
	DisplayName *string `json:"displayName,omitempty"`
	// 如不更新，则为空
	Description *string `json:"description,omitempty"`
}

// 更新数据源的输入
type UpdateDatasourceInput struct {
	// name, namespace用来确定资源
	// 规则: 不允许修改的。将原数据传递回来即可。
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	// 更新的的标签信息，这里涉及到增加或者删除标签，
	// 规则: 不允许修改的。将原数据传递回来即可。
	// 如果标签有任何改动，传递完整的label。
	// 例如之前的标齐是: abc:def 新增一个标签aa:bb, 那么传递 abc:def, aa:bb
	Labels      map[string]interface{} `json:"labels,omitempty"`
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 如不更新，则为空
	DisplayName *string `json:"displayName,omitempty"`
	// 如不更新，则为空
	Description *string `json:"description,omitempty"`
}

type UpdateEmbedderInput struct {
	// 模型服务资源名称（不可同名）
	Name string `json:"name"`
	// 模型服务创建命名空间
	Namespace string `json:"namespace"`
	// 模型服务资源标签
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 模型服务资源注释
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 模型服务资源展示名称作为显示，并提供编辑
	DisplayName *string `json:"displayName,omitempty"`
	// 模型服务资源描述
	Description *string `json:"description,omitempty"`
}

// 知识库更新的输入
type UpdateKnowledgeBaseInput struct {
	// 知识库资源名称（不可同名）
	Name string `json:"name"`
	// 知识库创建命名空间
	Namespace string `json:"namespace"`
	// 知识库资源标签
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 知识库资源注释
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 如不更新，则为空
	DisplayName *string `json:"displayName,omitempty"`
	// 如不更新，则为空
	Description *string `json:"description,omitempty"`
}

type UpdateModelInput struct {
	// 模型资源名称（不可同名）
	Name string `json:"name"`
	// 模型创建命名空间
	Namespace string `json:"namespace"`
	// 模型资标签
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 模型资源注释
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 模型资源展示名称作为显示，并提供编辑
	DisplayName *string `json:"displayName,omitempty"`
	// 模型资源描述
	Description *string `json:"description,omitempty"`
}

type UpdateVersionedDatasetInput struct {
	// 这个名字就是metadat.name, 根据name和namespace确定资源
	// name，namespac是不可以更新的。
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	// 更新的的标签信息，这里涉及到增加或者删除标签，
	// 所以，如果标签有任何改动，传递完整的label。
	// 例如之前的标齐是: abc:def 新增一个标签aa:bb, 那么传递 abc:def, aa:bb
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 传递方式同label
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	DisplayName *string                `json:"displayName,omitempty"`
	Description *string                `json:"description,omitempty"`
	// 更新，删除数据集版本中的文件，传递方式于label相同，完全传递。
	// 如果传递一个空的数组过去，认为是删除全部文件。
	FileGroups []*FileGroup `json:"fileGroups,omitempty"`
	// 修改数据集版本发布状态
	Released *int `json:"released,omitempty"`
}

// VersionedDataset
// 数据集的版本信息。
// 主要记录版本名字，数据的来源，以及文件的同步状态
type VersionedDataset struct {
	// 版本数据集id,为CR资源中的metadata.uid
	ID *string `json:"id,omitempty"`
	// 数据集名称, 这个应该是前端随机生成就可以，没有实际用途
	Name string `json:"name"`
	// 数据集所在的namespace，也是后续桶的名字
	Namespace string `json:"namespace"`
	// 一些用于标记，选择的的标签
	Labels map[string]interface{} `json:"labels,omitempty"`
	// 添加一些辅助性记录信息
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// 创建者，正查给你这个字段是不需要人写的，自动添加
	Creator *string `json:"creator,omitempty"`
	// 展示名字， 与metadat.name不一样，这个展示名字是可以用中文的
	DisplayName *string `json:"displayName,omitempty"`
	// 描述
	Description *string `json:"description,omitempty"`
	// 所属的数据集
	Dataset TypedObjectReference `json:"dataset"`
	// 更新时间, 这里更新指文件同步，或者数据处理完成后，做的更新操作的时间
	UpdateTimestamp   *time.Time `json:"updateTimestamp,omitempty"`
	CreationTimestamp time.Time  `json:"creationTimestamp"`
	// 数据集所包含的文件，对于文件需要支持过滤和分页
	Files PaginatedResult `json:"files"`
	// 版本名称
	Version string `json:"version"`
	// 该版本是否已经发布, 0是未发布，1是已经发布
	Released int `json:"released"`
	// 文件的同步状态, Processing或者'' 表示文件正在同步，Succeede 文件同步成功，Failed 存在文件同步失败
	SyncStatus *string `json:"syncStatus,omitempty"`
	// 数据处理状态，如果为空，表示还没有开始，processing 处理中，process_fail处理失败，process_complete处理完成
	DataProcessStatus *string `json:"dataProcessStatus,omitempty"`
}

func (VersionedDataset) IsPageNode() {}

type VersionedDatasetMutation struct {
	CreateVersionedDataset  VersionedDataset `json:"createVersionedDataset"`
	UpdateVersionedDataset  VersionedDataset `json:"updateVersionedDataset"`
	DeleteVersionedDatasets *string          `json:"deleteVersionedDatasets,omitempty"`
}

type VersionedDatasetQuery struct {
	GetVersionedDataset   VersionedDataset `json:"getVersionedDataset"`
	ListVersionedDatasets PaginatedResult  `json:"listVersionedDatasets"`
}

// 文件详情
// 描述: 文件在知识库中的详细状态
type Filedetail struct {
	// 文件路径
	Path string `json:"path"`
	// 文件处理的阶段
	// 规则: enum { Pending , Processing , Succeeded, Failed, Skipped}
	Phase string `json:"phase"`
}

// 文件组
// 规则: 属于同一个源(数据集)的文件要放在同一个filegroup中
// 规则: path直接读取文件里表中的文件路径即可
type Filegroup struct {
	// 源；目前仅支持版本数据集，即 Kind为 VersionedDataset
	Source *TypedObjectReference `json:"source,omitempty"`
	// 路径数组
	Path []string `json:"path,omitempty"`
}

// 文件组详情
// 描述: 文件组在知识库中的状态
type Filegroupdetail struct {
	// 源；目前仅支持版本数据集，即 Kind为 VersionedDataset
	Source *TypedObjectReference `json:"source,omitempty"`
	// 文件详情
	// 规则；数组。具体文件详情参考 filedetail描述
	Filedetails []*Filedetail `json:"filedetails,omitempty"`
}

// 源文件输入
type Filegroupinput struct {
	// 数据源字段
	Source TypedObjectReferenceInput `json:"source"`
	// 路径
	Path []string `json:"path,omitempty"`
}
