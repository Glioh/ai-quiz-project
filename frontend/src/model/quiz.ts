// Objective: Define the interfaces for the quiz model.


export interface Quiz {
    id: string;
    name: string;
    questions: QuizQuestion[];
}

export interface Player {
    id: string;
    name: string;
}

export interface QuizQuestion {
    id: string;
    question: string;
    choices: QuizChoice[];
}

export interface QuizChoice {
    id: string;
    choice: string;
    correct: boolean;
}
