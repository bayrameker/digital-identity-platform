import mongoose, { Document, Schema } from 'mongoose';

export interface INotification extends Document {
    userId: string;
    message: string;
    type: string;
    read: boolean;
    timestamp: Date;
}

const NotificationSchema: Schema = new Schema({
    userId: { type: String, required: true },
    message: { type: String, required: true },
    type: { type: String, enum: ['info', 'warning', 'alert'], default: 'info' },
    read: { type: Boolean, default: false },
    timestamp: { type: Date, default: Date.now },
});

export default mongoose.model<INotification>('Notification', NotificationSchema);
