import { Router } from 'express';
import { registerWebhook } from '../controllers/webhookController';
import { authMiddleware } from '../middlewares/authMiddleware';

const router = Router();

router.post('/register', authMiddleware, registerWebhook);

export default router;
