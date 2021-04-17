const cloud = require('wx-server-sdk')

cloud.init()

// 云函数入口函数
exports.main = async (event, context) => {
    try {
        const result = await cloud.openapi.updatableMessage.createActivityId()
        return result
        // {errCode:0,activityId:string	动态消息的 ID,expirationTime:number	activity_id 的过期时间戳。默认24小时后过期。}
    } catch (err) {
        throw err
    }
}