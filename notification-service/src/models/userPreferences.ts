import mongoose, { Document, Schema } from 'mongoose';

export interface IUserPreferences extends Document {
    userId: string;
    emailNotifications: boolean;
    pushNotifications: boolean;
    smsNotifications: boolean;
}

const UserPreferencesSchema: Schema = new Schema({
    userId: { type: String, required: true, unique: true },
    emailNotifications: { type: Boolean, default: true },
    pushNotifications: { type: Boolean, default: true },
    smsNotifications: { type: Boolean, default: false },
});

export default mongoose.model<IUserPreferences>('UserPreferences', UserPreferencesSchema);
