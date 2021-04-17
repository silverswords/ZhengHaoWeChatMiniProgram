import {
  addPicture,
  queryPList,
  removePicture,
  updatePicture,
} from "@/services/api";
import router from "umi/router";

export default {
  namespace: "list",
  state: {
    data: {},
    list: []
  },

  effects: {
    *queryPictureList({ payload }, { call, put }) {
      const response = yield call(queryPList, payload);
      if (response.status !== 200) {
        return
      }
      if (response.Projects === null){
        response.Projects = []
      }
      yield put({
        type: "queryList",
        payload: response.Projects
      });
    },
    *createorupdate({ payload }, { call, put }) {
      let callback;
      if(payload.ProjectId) {
        callback = updatePicture;
      } else {
        callback = addPicture;
      }
      const response = yield call(callback, payload);
      if(response === undefined) {
        router.push("/exception/500")
      }
      yield put({
        type: 'queryPictureList',
        payload: response,
      })
    },
    *delete({ payload, callback }, { call, put }) {
      const response = yield call(removePicture, payload);
      //更新删除后数据
      if (response.status === 200) {
        const response = yield call(queryPList, payload);
        yield put({
          type: "queryPictureList",
          payload: response
        });
      }
    },
  },

  reducers: {
    queryList(state, { payload }) {
      return {
        list: payload
      };
    },
  }
};
