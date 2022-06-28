import axios from 'axios';
import Cookies from "js-cookie";

class Api {
    
    // 起点榜单数据API
    QIDIAN_SORT = '/api/qidian/';

    // 对axios的一个封装
    send(obj){
        return axios({
            method: obj.method?obj.method:"POST",
            url: obj.url,
            data: obj.data?obj.data:{},
            headers: {
                'Content-Type': 'application/json',
                'Authorization': Cookies.get("token")?Cookies.get("token"):""
            },
            params: obj.params
        })
    }
}

export default Api;