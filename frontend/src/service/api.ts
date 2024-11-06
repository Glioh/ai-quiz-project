import type { Quiz } from "../model/quiz";

export class ApiService {
    private baseUrl: string;

    constructor() {
        // MAY HAVE TO ADD WS
        this.baseUrl = 'http://localhost:3000';
    }

    async deleteQuiz(quizId: string): Promise<void> {
        const response = await fetch(`${this.baseUrl}/api/quizzes/${quizId}`, {
            method: 'DELETE',
        });
        
        if (!response.ok) {
            throw new Error('Failed to delete quiz');
        }
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

    async generateAIQuiz(name: string, prompt: string, numQuestions: number = 10) {
        try {
            console.log('Sending request to:', `${this.baseUrl}/api/quiz/generate`);
            console.log('Request payload:', { name, prompt, numQuestions });
    
            const response = await fetch(`${this.baseUrl}/api/quiz/generate`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ name, prompt, numQuestions }),
            });
    
            console.log('Response status:', response.status);
    
            if (!response.ok) {
                const errorText = await response.text();
                console.error('Error response:', errorText);
                throw new Error(`HTTP error! status: ${response.status}, message: ${errorText}`);
            }
    
            const data = await response.json();
            console.log('Response data:', data);
            return data;
        } catch (error) {
            console.error('API Error:', error);
            throw error;
        }
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