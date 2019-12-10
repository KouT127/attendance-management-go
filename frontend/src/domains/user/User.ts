export class User extends Document {
    id: string;
    username: string;
    email: string;
    imageUrl: string;
    shouldEdit: boolean;

    constructor(id: string, username: string, email: string, imageUrl: string, shouldEdit = true) {
        super();
        this.id = id;
        this.username = username;
        this.email = email;
        this.imageUrl = imageUrl;
        this.shouldEdit = shouldEdit;
    }

    static initializeUser(userId: string): User {
        return new User(userId, '', '', '', true)
    }

    toJson() {
        return {
            user_name: this.username,
            email: this.email,
            imageUrl: this.imageUrl,
        }
    }
}
