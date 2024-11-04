import type { Quiz } from "../model/quiz";

export class ApiService {
    private baseUrl: string;

    constructor() {
        // MAY HAVE TO ADD WS
        this.baseUrl = 'http://localhost:3000';
    }

    async getQuizById(id: string): Promise<Quiz | null> {
        let response = await fetch(`http://localhost:3000/api/quizzes/${id}`);
        if (!response.ok) {
            return null;
        }

        let json = await response.json();
        return json;
    }

    async createQuiz(name: string): Promise<Quiz> {
        const response = await fetch(`${this.baseUrl}/api/quizzes`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name }),
        });
        
        return await response.json();
    }

    async getQuizzes(): Promise<Quiz[]> {
        let response = await fetch("http://localhost:3000/api/quizzes");
        if (!response.ok) {
            alert("Failed to fetch quizzes!");
            return [];
        }

        let json = await response.json();
        return json;
    }

    async saveQuiz(quizId: string, quiz: Quiz) {
        let response = await fetch(`http://localhost:3000/api/quizzes/${quizId}`, {
            method: "PUT",
            body: JSON.stringify(quiz),
            headers: {
                "Content-Type": "application/json"
            }
        });

        if (!response.ok) {
            alert("Failed to save quiz!");
            return;
        }
    }
}

export const apiService = new ApiService();