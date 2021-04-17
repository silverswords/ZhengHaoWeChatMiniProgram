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
@connect(({ rank, loading }) => ({
  rank,
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
    imageUrl: '',
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
      type: 'rank/ranking',
      payload: {
        count: 5,
      },
    });
  }

  render() {
    const {
      rank: { rank },
      loading,
    } = this.props;
    const {
      form: { getFieldDecorator },
    } = this.props;
    const { fileList, previewVisible, previewVisibleQR, visible, done, current, qrFileList } = this.state;
    const previewImage = current.Path
    const previewImageQR = current.QRCode


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

    const ListContent = ({ data: { Number, Path, Points, UserId, UserName } }) => (
      <div className={styles.listContent}>
        <div className={styles.listContentItem}>
          <span>积分数</span>
          <p>{Points}</p>
        </div>
      </div>
    );


    return (
      <PageHeaderWrapper>
        <div className={styles.standardList}>
          <Card
            className={styles.listCard}
            bordered={false}
            title="排行榜"
            style={{ marginTop: 24 }}
            bodyStyle={{ padding: '0 32px 40px 32px' }}
          >
            <List
              size="large"
              rowKey="id"
              loading={loading}
              // pagination={paginationProps}
              dataSource={rank}
              renderItem={item => (
                <List.Item>
                  <List.Item.Meta
                    avatar={<Avatar src={item.Path} shape="square" size="large" />}
                    title={<p>{item.UserName}</p>}
                  />
                  <ListContent data={item} />
                </List.Item>
              )}
            />
          </Card>
        </div>

      </PageHeaderWrapper>
    );
  }
}

export default BasicList;
