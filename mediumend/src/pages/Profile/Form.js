import React, { PureComponent } from 'react';
import { findDOMNode } from 'react-dom';
import moment from 'moment';
import { connect } from 'dva';
import {
  List,
  Card,
  Row,
  Col,
  Radio,
  Input,
  InputNumber,
  Progress,
  Button,
  Icon,
  message,
  Dropdown,
  Menu,
  Avatar,
  Modal,
  Form,
  DatePicker,
  Select,
  Upload,
} from 'antd';
import { UploadOutlined } from '@ant-design/icons';


import PageHeaderWrapper from '@/components/PageHeaderWrapper';
import Result from '@/components/Result';

import styles from './Form.less';

const FormItem = Form.Item;
const RadioButton = Radio.Button;
const RadioGroup = Radio.Group;
const SelectOption = Select.Option;
const { Search, TextArea } = Input;
var Minio = require('minio');
var minioClient = new Minio.Client({
  endPoint: 'xcx.zhenghaodichan.com',
  port: 9000,
  useSSL: true,
  accessKey: 'minioadmin',
  secretKey: 'minioadmin'
});
const bucketName = 'test'
@connect(({ list, loading }) => ({
  list,
  loading: loading.models.list,
}))
@Form.create()
class BasicList extends PureComponent {
  state = {
    visible: false,
    done: false,
    fileList: [],
    qrFileList: [],
    previewVisible: false,
    previewVisibleQR: false,
    imageUrls: [],
    qrCodeUrl: '',
    current: {},
    delete: false,
  };

  formLayout = {
    labelCol: { span: 7 },
    wrapperCol: { span: 13 },
  };

  componentDidMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'list/queryPictureList',
      payload: {
        count: 5,
      },
    });
  }

  showModal = () => {
    this.setState({
      visible: true,
      current: {},
      fileList: [],
      qrFileList: [],
    });
  };

  showEditModal = item => {
    if (item.Path != ' ') {
      console.log(item)
      let fileList = []
      let imageUrls = []
      if (item.PathOne !== ' ') {
        imageUrls.push({
          uid: '1',
          imageUrl: item.PathOne,
        })
        fileList.push({
          uid: '1',
          name: 'image.png',
          status: 'done',
          url: item.PathOne,
        })
      }
      if (item.PathTwo !== ' ') {
        imageUrls.push({
          uid: '2',
          imageUrl: item.PathTwo,
        })
        fileList.push({
          uid: '2',
          name: 'image.png',
          status: 'done',
          url: item.PathTwo,
        })
      }
      if (item.PathThree !== ' ') {
        imageUrls.push({
          uid: '3',
          imageUrl: item.PathThree,
        })
        fileList.push({
          uid: '3',
          name: 'image.png',
          status: 'done',
          url: item.PathThree,
        })
      }
      if (item.PathFour !== ' ') {
        imageUrls.push({
          uid: '4',
          imageUrl: item.PathFour,
        })
        fileList.push({
          uid: '4',
          name: 'image.png',
          status: 'done',
          url: item.PathFour,
        })
      }
      if (item.PathFive !== ' ') {
        imageUrls.push({
          uid: '5',
          imageUrl: item.PathFive,
        })
        fileList.push({
          uid: '5',
          name: 'image.png',
          status: 'done',
          url: item.PathFive,
        })
      }
      if (item.PathSix !== ' ') {
        imageUrls.push({
          uid: '6',
          imageUrl: item.PathSix,
        })
        fileList.push({
          uid: '6',
          name: 'image.png',
          status: 'done',
          url: item.PathSix,
        })
      }
      if (item.PathSeven !== ' ') {
        imageUrls.push({
          uid: '7',
          imageUrl: item.PathSeven,
        })
        fileList.push({
          uid: '7',
          name: 'image.png',
          status: 'done',
          url: item.PathSeven,
        })
      }
      if (item.PathEight !== ' ') {
        imageUrls.push({
          uid: '8',
          imageUrl: item.PathEight,
        })
        fileList.push({
          uid: '8',
          name: 'image.png',
          status: 'done',
          url: item.PathEight,
        })
      }
      if (item.PathNine !== ' ') {
        imageUrls.push({
          uid: '9',
          imageUrl: item.PathNine,
        })
        fileList.push({
          uid: '9',
          name: 'image.png',
          status: 'done',
          url: item.PathNine,
        })
      }

      this.setState({
        visible: true,
        current: item,
        fileList: fileList,
        imageUrls: imageUrls,
        qrFileList: [{
          uid: '-1',
          name: 'image.png',
          status: 'done',
          url: item.QRCode,
        },],
      });
    }
    this.setState({
      visible: true,
      current: item,
    });
  };

  handleDone = () => {
    setTimeout(() => this.addBtn.blur(), 0);
    this.setState({
      done: false,
      visible: false,
    });
  };

  handleCancel = () => {
    setTimeout(() => this.addBtn.blur(), 0);
    this.setState({
      visible: false,
      imageUrls: [],
    });
  };

  handleSubmit = e => {
    e.preventDefault();
    const { dispatch, form } = this.props;
    const { qrCodeUrl, imageUrls, current } = this.state;
    const ProjectId = current ? current.ProjectID : '';

    form.validateFields((err, fieldsValue) => {
      if (err) return;
      this.setState({
        done: true,
        imageUrls: [],
        qrCodeUrl: '',
      });
      let qr = qrCodeUrl === '' ? current.QRCode : qrCodeUrl
      for (let i = 0; i < 9; i++) {
        if (imageUrls[i] === undefined) {
          imageUrls[i] = { uid: " ", imageUrl: " " }
        }
      }
      dispatch({
        type: 'list/createorupdate',
        payload: {
          ProjectId, ...fieldsValue,
          pathOne: imageUrls[0].imageUrl,
          pathTwo: imageUrls[1].imageUrl,
          pathThree: imageUrls[2].imageUrl,
          pathFour: imageUrls[3].imageUrl,
          pathFive: imageUrls[4].imageUrl,
          pathSix: imageUrls[5].imageUrl,
          pathSeven: imageUrls[6].imageUrl,
          pathEight: imageUrls[7].imageUrl,
          pathNine: imageUrls[8].imageUrl, qrCode: qr,
        },
      });
    });
  };

  deleteItem = ProjectId => {
    const { dispatch } = this.props;
    dispatch({
      type: 'list/delete',
      payload: { ProjectId },
    });
  };


  //上传图片
  handleChange = (info) => {
    const { dispatch } = this.props;
    const { imageUrls } = this.state;

    let imageUrlsTemp = imageUrls
    let fileList = info.fileList;
    console.log(fileList, "fileList", fileList[4], info.file)

    this.setState({ fileList });

    function findFunc(uid) {
      return (picture) => {
        return picture.uid === uid
      }
    }
    if (imageUrlsTemp.find(findFunc(info.file.uid))) {
      // remove the image
      console.log("remove the picture")
      imageUrlsTemp = imageUrlsTemp.filter(({ uid }) => uid !== info.file.uid)
      console.log(imageUrlsTemp, info.file.uid)
      this.setState({
        imageUrls: imageUrlsTemp
      })
      // do the remove
      return
    }
    let self = this;
    minioClient.presignedPutObject(bucketName, info.file.name, 60, function (err, presignedUrl) {
      if (err) return console.log(err)
      console.log(presignedUrl, info.file)
      fetch(presignedUrl, {
        method: 'PUT',
        body: info.file
      }).then(() => {
        imageUrlsTemp.push({ uid: info.file.uid, imageUrl: "https://xcx.zhenghaodichan.com:9000/" + bucketName + "/" + info.file.name })
        self.setState({
          imageUrls: imageUrlsTemp
        })
        console.log(imageUrlsTemp)
      }
      )
    })
  }

  handleChangeQR = (info) => {
    const { dispatch } = this.props;
    let fileList = info.fileList;
    this.setState({ qrFileList: fileList });
    if (fileList.length === 0) {
      // remove the image
      return
    }
    let self = this;
    minioClient.presignedPutObject(bucketName, info.file.name, 60, function (err, presignedUrl) {
      if (err) return console.log(err)
      console.log(presignedUrl, info.file)
      fetch(presignedUrl, {
        method: 'PUT',
        body: info.file
      }).then(() => {
        self.setState({
          qrCodeUrl: "https://xcx.zhenghaodichan.com:9000/" + bucketName + "/" + info.file.name
        })
      }
      )
    })
  }

  beforeUpload() {
    return false
  }

  handlePreviewCancel = () => this.setState({ previewVisible: false })

  handlePreviewQR = (file) => {
    this.setState({
      previewImageQR: file.url || file.thumbUrl,
      previewVisibleQR: true,
    });
  }

  // handlePreview = (file) => {
  //   console.log(file.url)
  //   this.setState({
  //     previewImage: file.url,
  //     previewVisible: true,
  //   });
  // }

  render() {
    const {
      list: { list },
      loading,
    } = this.props;
    const {
      form: { getFieldDecorator },
    } = this.props;
    const { fileList, previewVisible, previewVisibleQR, visible, done, current, qrFileList } = this.state;
    const previewImage = current.Path
    const previewImageQR = current.QRCode

    const editAndDelete = (key, currentItem) => {
      if (key === 'edit') this.showEditModal(currentItem);
      else if (key === 'delete') {
        Modal.confirm({
          title: '删除项目',
          content: '确定删除该项目吗？',
          okText: '确认',
          cancelText: '取消',
          onOk: () => this.deleteItem(currentItem.ProjectID),
        });
      }
      this.setState({ delete: true })
    };

    const uploadButton = (
      <div>
        <Icon type={this.state.loading ? 'loading' : 'plus'} />
        <div>Upload</div>
      </div>
    )
    const modalFooter = done
      ? { footer: null, onCancel: this.handleDone }
      : { okText: '保存', onOk: this.handleSubmit, onCancel: this.handleCancel };

    const Info = ({ title, value, bordered }) => (
      <div className={styles.headerInfo}>
        <span>{title}</span>
        <p>{value}</p>
        {bordered && <em />}
      </div>
    );

    const paginationProps = {
      showSizeChanger: true,
      showQuickJumper: true,
      pageSize: 5,
      total: 50,
    };

    const ListContent = ({ data: { AddPoints, Introduction, Path, ProjectID, QRCode, Rule } }) => (
      <div className={styles.listContent}>
        <div className={styles.listContentItem}>
          <span>项目介绍</span>
          <p>{Introduction}</p>
        </div>
        <div className={styles.listContentItem}>
          <span>每次转发积分数</span>
          <p>{AddPoints}</p>
        </div>
        <div className={styles.listContentItem}>
          <p>转发规则</p>
          <p>{Rule}</p>
        </div>
      </div>
    );

    const MoreBtn = props => (
      <Dropdown
        overlay={
          <Menu onClick={({ key }) => editAndDelete(key, props.current)}>
            <Menu.Item key="edit">编辑</Menu.Item>
            <Menu.Item key="delete">删除</Menu.Item>
          </Menu>
        }
      >
        <a>
          更多 <Icon type="down" />
        </a>
      </Dropdown>
    );

    const getModalContent = () => {
      if (done) {
        return (
          <Result
            type="success"
            title="操作成功"
            description="一系列的信息描述，很短同样也可以带标点。"
            actions={
              <Button type="primary" onClick={this.handleDone}>
                知道了
              </Button>
            }
            className={styles.formResult}
          />
        );
      }
      return (
        <Form onSubmit={this.handleSubmit}>
          <FormItem label="项目名称" {...this.formLayout}>
            {getFieldDecorator('projectName', {
              rules: [{ required: true, message: '请输入项目名称' }],
              initialValue: current.ProjectName,
            })(<Input placeholder="请输入" />)}
          </FormItem>
          <FormItem label="上传项目图" {...this.formLayout}>
            {getFieldDecorator("path", {
              initialValue: current.PathOne || '',
              rules: [{ required: true, message: "请上传项目预览图" }]
            })(
              <div>
                <Upload
                  name="ImagePath"
                  listType="picture-card"
                  fileList={fileList}
                  beforeUpload={this.beforeUpload}
                  // onPreview={this.handlePreview}
                  onChange={this.handleChange}
                  accept="image/*"
                >
                  {fileList.length >= 9 ? null : uploadButton}
                </Upload>
                <Modal visible={previewVisible} footer={null} onCancel={this.handlePreviewCancel}>
                  <img alt="image" style={{ width: '100%' }} src={previewImage} />
                </Modal>
              </div>
            )}
          </FormItem>
          <FormItem label="加砖数量" {...this.formLayout}>
            {getFieldDecorator('addPoints', {
              rules: [{ required: true, message: '请输入加砖数量' }],
              initialValue: current.AddPoints,
            })(<InputNumber placeholder="请输入" />)}
          </FormItem>
          <FormItem {...this.formLayout} label="项目简介">
            {getFieldDecorator('introduction', {
              rules: [{ message: '请输入至少五个字符的项目简介！', min: 5 }],
              initialValue: current.Introduction,
            })(<TextArea rows={4} placeholder="请输入至少五个字符" />)}
          </FormItem>
          <FormItem {...this.formLayout} label="项目规则">
            {getFieldDecorator('rule', {
              rules: [{ message: '请输入至少三个字符的项目规则！', min: 3 }],
              initialValue: current.Rule,
            })(<TextArea rows={4} placeholder="请输入至少三个字符" />)}
          </FormItem>
          <FormItem label="上传二维码" {...this.formLayout}>
            {getFieldDecorator("qrCode", {
              initialValue: current.QRCode || '',
              rules: [{ required: true, message: "请上传二维码" }]
            })(
              <div>
                <Upload
                  name="qrCode"
                  listType="picture-card"
                  fileList={qrFileList}
                  beforeUpload={this.beforeUpload}
                  onChange={this.handleChangeQR}
                  accept="image/*"
                >
                  {qrFileList.length >= 1 ? null : uploadButton}
                </Upload>
                <Modal visible={previewVisibleQR} footer={null} onCancel={this.handlePreviewCancel}>
                  <img alt="image" style={{ width: '100%' }} src={previewImageQR} />
                </Modal>
              </div>
            )}
          </FormItem>
        </Form>

      );
    };
    return (
      <PageHeaderWrapper>
        <div className={styles.standardList}>
          <Card
            className={styles.listCard}
            bordered={false}
            title="项目列表"
            style={{ marginTop: 24 }}
            bodyStyle={{ padding: '0 32px 40px 32px' }}
          >
            <Button
              type="dashed"
              style={{ width: '100%', marginBottom: 8 }}
              icon="plus"
              onClick={this.showModal}
              ref={component => {
                /* eslint-disable */
                this.addBtn = findDOMNode(component);
                /* eslint-enable */
              }}
            >
              添加
            </Button>
            <List
              size="large"
              rowKey="id"
              loading={loading}
              // pagination={paginationProps}
              dataSource={list}
              renderItem={item => (
                <List.Item
                  actions={[
                    <a
                      onClick={e => {
                        e.preventDefault();
                        this.showEditModal(item);
                      }}
                    >
                      编辑
                    </a>,
                    <MoreBtn current={item} />,
                  ]}
                >
                  <List.Item.Meta
                    avatar={<Avatar src={item.PathOne} shape="square" size="large" />}
                    title={<p>{item.ProjectName}</p>}
                  />
                  <ListContent data={item} />
                </List.Item>
              )}
            />
          </Card>
        </div>
        <Modal
          title={done ? null : `任务`}
          className={styles.standardListForm}
          width={640}
          bodyStyle={done ? { padding: '72px 0' } : { padding: '28px 0 0' }}
          destroyOnClose
          visible={visible}
          {...modalFooter}
        >
          {getModalContent()}
        </Modal>
      </PageHeaderWrapper>
    );
  }
}

export default BasicList;
