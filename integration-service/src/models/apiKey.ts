import mongoose, { Document, Schema } from 'mongoose';

export interface IApiKey extends Document {
    userId: string;
    key: string;
    createdAt: Date;
    expiresAt: Date;
}

const ApiKeySchema: Schema = new Schema({
    userId: { type: String, required: true },
    key: { type: String, required: true, unique: true },
    createdAt: { type: Date, default: Date.now },
    expiresAt: { type: Date, required: true },
});

export default mongoose.model<IApiKey>('ApiKey', ApiKeySchema);
