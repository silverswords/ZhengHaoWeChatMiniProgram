import { routerRedux } from "dva/router";
import { stringify } from "qs";
import { fakeAccountLogin, getFakeCaptcha } from "@/services/api";
import { setAuthority } from "@/utils/authority";
import { getPageQuery } from "@/utils/utils";
import { reloadAuthorized } from "@/utils/Authorized";

export default {
  namespace: "login",

  state: {
    status: undefined
  },

  effects: {
    *login({ payload }, { call, put }) {
      const response = yield call(fakeAccountLogin, payload);
      yield put({
        type: "changeLoginStatus",
        payload: response,
      });
      // const { status } = response;

      // if ((status !== undefined) && (status === ok)){
      //   setToken(response.data)
      // }
      console.log(response, "response")
      // if (response.currentAuthority != '') {
      //   window.location.href = "http://localhost:8000/article/article-list";
      //   return;
      // }
      
      if (response.status === 200) {
        window.location.href = "https://xcx.zhenghaodichan.com/profile";
        // window.location.href = "http://localhost:8001/profile";
        return;
      }
    }

    // *getCaptcha({ payload }, { call }) {
    //   yield call(getFakeCaptcha, payload);
    // },

    // *logout(_, { put }) {
    //   yield put({
    //     type: "changeLoginStatus",
    //     payload: {
    //       status: false,
    //       currentAuthority: "guest"
    //     }
    //   });
    //   reloadAuthorized();
    //   yield put(
    //     routerRedux.push({
    //       pathname: "/user/login",
    //       search: stringify({
    //         redirect: window.location.href
    //       })
    //     })
    //   );
    // }
  },

  reducers: {
    // changeLoginStatus(state, { payload }) {
    //   //setAuthority(payload.currentAuthority);
    //   return {
    //     ...state,
    //     status: payload.status,
    //   };
    // }
    changeLoginStatus(state, { payload }) {
      setAuthority(payload.currentAuthority);
      return {
        ...state,
        status: payload.status,
        type: payload.type,
      };
    },
  }
};
