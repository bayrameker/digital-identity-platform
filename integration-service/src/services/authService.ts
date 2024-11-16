import bcrypt from 'bcryptjs';
import jwt from 'jsonwebtoken';
import User, { IUser } from '../models/user';
import config from '../utils/config';

export const registerUser = async (username: string, password: string) => {
    const hashedPassword = await bcrypt.hash(password, 10);
    const user = new User({ username, password: hashedPassword });
    await user.save();
    return user;
};

export const loginUser = async (username: string, password: string) => {
    const user = await User.findOne({ username });
    if (!user) {
        throw new Error('Invalid credentials');
    }
    const isMatch = await bcrypt.compare(password, user.password);
    if (!isMatch) {
        throw new Error('Invalid credentials');
    }
    const token = jwt.sign({ userId: user._id }, config.get('jwtSecret'), { expiresIn: '1h' });
    return token;
};
