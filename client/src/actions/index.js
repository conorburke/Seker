import axios from 'axios';

import { FETCH_TOOLS } from './types';
import url from '../utils';

export const fetchTools = () => {
    //return function to have redux-thunk use dispact function
    //normally would just return object, but thunk will see function
    //and apply middleware
    return function(dispatch) {
        axios.get(`${url.api}/tools`)
            .then(res => dispatch({type: FETCH_TOOLS, payload: res.data}));
    }
};