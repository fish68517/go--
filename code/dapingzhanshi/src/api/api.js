import {get} from './http';
export const getBigScreenData = param => get('/bigScreenData/getLatestScreenData',param);