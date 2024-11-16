import { Request, Response } from 'express';
import { fetchThirdPartyData, postThirdPartyData } from '../services/apiService';

export const getData = async (req: Request, res: Response) => {
    try {
        const data = await fetchThirdPartyData('https://api.thirdparty.com/data', req.query);
        res.status(200).json(data);
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
};

export const postData = async (req: Request, res: Response) => {
    try {
        const data = await postThirdPartyData('https://api.thirdparty.com/data', req.body);
        res.status(200).json(data);
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
};
