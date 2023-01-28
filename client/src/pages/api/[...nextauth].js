import NextAuth from "next-auth";
import GoogleProvider from "next-auth/providers/google";
import GithubProvider from "next-auth/providers/github";
import CredentialsProvider from "next-auth/providers/credentials";
import axios from "axios";

const providers = [
    // Google Provider
    GoogleProvider({
        clientId: process.env.GOOGLE_ID,
        clientSecret: process.env.GOOGLE_SECRET,
    }),
    GithubProvider({
        clientId: process.env.GITHUB_ID,
        clientSecret: process.env.GITHUB_SECRET,
    }),
    CredentialsProvider({
        name: "Credentials",
        async authorize(credentials, req) {
            const user = await axios.post("/api/users/signup", {
                credentials,
            });

            console.log(user);

            return user.data;
        },
    }),
];

const callbacks = {
    async jwt(token, user) {
        console.log(user);
        if (user) {
            token.accessToken = user.data.token;
        }

        return token;
    },

    async session(session, token) {
        session.accessToken = token.accessToken;
        return session;
    },
};

const options = {
    providers,
    callbacks,
    pages: {
        error: "/login", // Changing the error redirect page to our custom login page
    },
};

export default (req, res) => NextAuth(req, res, options);
