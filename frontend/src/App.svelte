<script lang="ts">
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'
  import Counter from './lib/Counter.svelte'
  import Button from './lib/Button.svelte'
  import QuizCard from './lib/QuizCard.svelte'
  import type { Quiz, QuizQuestion } from './model/quiz';
  import { NetService, PacketTypes, type ChangeGameStatePacket } from './service/net';

  let quizzes: {_id : string, name: string}[] = [];
  
  let currentQuestion: QuizQuestion | null = null;
  let state = -1;
  let netService = new NetService();
  netService.connect();
  netService.onPacket((packet: any) => {
    console.log(packet);
    switch(packet.id){
      case 2:{
        currentQuestion = packet.question;
        break;
      }

      case PacketTypes.ChangeGameState: {
        let data = packet as ChangeGameStatePacket
        state = data.state;
        break;
      }
    }
  });


  async function getQuizzes() {
    // fetch data from the server
    let response = await fetch('http://localhost:3000/api/quizzes');
    // check if the response is ok
    if(!response.ok) {
      alert("Failed!");
      return;
    }

      let json = await response.json();
      quizzes = json;
      console.log(json);
  }

  let code ="";
  let msg = "";
  let name = "";

  function connect() {
    netService.sendPacket({
      id: 0,
      code: code,
      name: name
    });
  }

  function hostQuiz(quiz: any) {
    netService.sendPacket({
      id: 1,
      quizId: quiz.id
    });
  }
</script>

{#if state == -1}
  <Button on:click={getQuizzes}>Get Quizzes</Button>
  Message: {msg}

  <div>
    {#each quizzes as quiz}
      <QuizCard on:host={() => hostQuiz(quiz)} quiz={quiz} />
    {/each}
  </div>

  <input bind:value={code} class="border" type="text" placeholder="Game Code" />
  <input bind:value={name} class="border" type="text" placeholder="Name" />
  <Button on:click={connect}>Join Game</Button>

  {#if currentQuestion != null}
    <h2 class="text-4x1 font-bold mt-8">{currentQuestion.name}</h2>
    <div class="flex">
    {#each currentQuestion.choices as choice}
      <div class="flex-1 bg-blue-400 text-center font-bold text-2x1 text-white justify-center">
        {choice.name}
      </div>
    {/each}
    </div>
  {/if}
{:else if state == 0}
    <p>Lobby State</p>
{/if}
