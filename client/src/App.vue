<script lang="ts">
type Game = {
  length: number;
};

export default {
  data: () => ({
    wordLength: 0,
    attempts: 5,
    currentRow: 1,
    currentChars: {} as any,
  }),
  created() {
    this.fetchWordLength();
  },
  methods: {
    async fetchWordLength() {
      const game = (await (
        await fetch("http://localhost:8080/new_game")
      ).json()) as unknown as Game;

      this.wordLength = game.length;
    },
    handleChange(n: string, e: Event) {
      this.currentChars[n] = e.target.value;

      let nInt = Number(n);

      if (e.target.value === "" && !n.endsWith("1")) {
        nInt -= 1;
      } else if (
        e.target.value !== "" &&
        !n.endsWith(String(this.wordLength))
      ) {
        nInt += 1;
      } else {
        return;
      }

      const nextInputRef = this.$refs[`${nInt}`];

      nextInputRef[0].focus();
    },
    isLetter(e: KeyboardEvent) {
      let char = String.fromCharCode(e.keyCode); // Get the character
      if (/^[A-Za-z]+$/.test(char)) return true; // Match with regex
      else e.preventDefault(); // If not match, don't add to input text
    },
    async handleSubmit(e: Event) {
      this.currentChars;
    },
  },
};
</script>

<template>
  <header>
    <h1>Termo</h1>
    <h2>A palavra tem {{ wordLength }} caracteres</h2>
  </header>

  <main>
    <template v-for="n in attempts">
      <form @submit.prevent="handleSubmit">
        <template v-for="w in wordLength">
          <input
            type="text"
            :disabled="n !== currentRow"
            maxlength="1"
            class="char"
            :name="`${n}${w}-char`"
            spellcheck="false"
            autocomplete="off"
            :value="currentChars[`${n}${w}`]"
            :id="`${n}${w}`"
            :ref="`${n}${w}`"
            @keypress="isLetter"
            @input="handleChange(`${n}${w}`, $event)"
          />
        </template>
      </form>
    </template>
  </main>
</template>

<style scoped>
header {
  color: aliceblue;
  display: flex;
  flex-direction: column;
  align-items: center;
}

main {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.char {
  width: 60px;
  height: 100px;
  margin: 10px;
  text-align: center;
  font-size: 2rem;
}
</style>
