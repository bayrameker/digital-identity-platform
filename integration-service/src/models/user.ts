import mongoose, { Document, Schema } from 'mongoose';

export interface IUser extends Document {
    username: string;
    password: string;
    apiKeys: string[];
}

const UserSchema: Schema = new Schema({
    username: { type: String, required: true, unique: true },
    password: { type: String, required: true },
    apiKeys: [{ type: String }],
});

export default mongoose.model<IUser>('User', UserSchema);
