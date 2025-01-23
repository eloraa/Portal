import { PUBLIC_API_URL } from '$env/static/public';

interface Member {
	userId: string;
	username: string;
	avatarId: string;
}

interface ChatMessage {
	fromUserId: string;
	message: string;
	timestamp: number;
}

export class WebRTCChat {
	private ws: WebSocket;
	private peerConnection?: RTCPeerConnection;
	private dataChannel?: RTCDataChannel;
	private messageCallback?: (message: ChatMessage) => void;
	private members: Member[] = [];
	private statusCallback?: (members: Member[]) => void;
	private signalCallback?: (signal: string) => void;

	constructor(
		private roomId: string,
		private userId: string,
		private username: string,
		private avatarId: string
	) {
		const baseUrl = PUBLIC_API_URL.replace('http://', 'ws://');
		this.ws = new WebSocket(`${baseUrl}/ws`);
		this.setupWebSocket();
	}

	private setupWebSocket() {
		this.ws.onopen = () => {
			this.ws.send(
				JSON.stringify({
					type: 'join_room',
					roomId: this.roomId,
					userId: this.userId,
					payload: {
						username: this.username,
						avatarId: this.avatarId
					}
				})
			);
		};

		this.ws.onmessage = async (event) => {
			const message = JSON.parse(event.data);
			switch (message.type) {
				case 'room_joined':
					this.members = message.payload.members;
					break;
				case 'user_joined': {
					const newMember = {
						userId: message.userId,
						username: message.payload.username,
						avatarId: message.payload.avatarId
					};
					this.members.push(newMember);
					break;
				}
				case 'user_left':
					this.members = this.members.filter((m) => m.userId !== message.userId);
					break;
			}
		};
	}
}
