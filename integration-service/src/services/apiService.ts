import axios from 'axios';

export const fetchThirdPartyData = async (endpoint: string, params: any) => {
    const response = await axios.get(endpoint, { params });
    return response.data;
};

export const postThirdPartyData = async (endpoint: string, data: any) => {
    const response = await axios.post(endpoint, data);
    return response.data;
};
