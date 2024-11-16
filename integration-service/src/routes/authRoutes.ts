import { Router } from 'express';
import { register, login } from '../controllers/authController';
import { body } from 'express-validator';
import { validate } from '../utils/validator';

const router = Router();

router.post(
    '/register',
    [
        body('username').isString().notEmpty(),
        body('password').isString().isLength({ min: 6 }),
        validate,
    ],
    register
);

router.post(
    '/login',
    [
        body('username').isString().notEmpty(),
        body('password').isString().notEmpty(),
        validate,
    ],
    login
);

export default router;
