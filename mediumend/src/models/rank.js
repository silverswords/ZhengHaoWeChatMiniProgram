import {
    queryRankList,
  } from "@/services/api";
  import router from "umi/router";
  
  export default {
    namespace: "rank",
    state: {
      data: {},
      rank: []
    },
  
    effects: {
      *ranking({ payload }, { call, put }) {
        const response = yield call(queryRankList, payload);
        if (response.status !== 200) {
          return
        }
        console.log(response)
        if (response.rankings === null){
          response.rankings = []
        }
        yield put({
          type: "queryRank",
          payload: response.rankings
        });
      },
    },
  
    reducers: {
      queryRank(state, { payload }) {
        return {
          rank: payload
        };
      },
    }
  };
  