import { Router } from 'express';
import { getData, postData } from '../controllers/apiController';
import { authMiddleware } from '../middlewares/authMiddleware';

const router = Router();

router.get('/data', authMiddleware, getData);
router.post('/data', authMiddleware, postData);

export default router;
