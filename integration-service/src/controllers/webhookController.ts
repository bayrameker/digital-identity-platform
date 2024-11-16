import { Request, Response } from 'express';
import webhookService from '../services/webhookService';

export const registerWebhook = async (req: Request, res: Response) => {
    try {
        const { url, event } = req.body;
        await webhookService.registerWebhook(url, event);
        res.status(201).json({ message: 'Webhook registered' });
    } catch (error) {
        res.status(400).json({ error: error.message });
    }
};
