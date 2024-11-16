import mongoose, { Document, Schema } from 'mongoose';

export interface IOAuthClient extends Document {
    clientId: string;
    clientSecret: string;
    redirectUris: string[];
    grants: string[];
}

const OAuthClientSchema: Schema = new Schema({
    clientId: { type: String, required: true, unique: true },
    clientSecret: { type: String, required: true },
    redirectUris: [{ type: String }],
    grants: [{ type: String }],
});

export default mongoose.model<IOAuthClient>('OAuthClient', OAuthClientSchema);
