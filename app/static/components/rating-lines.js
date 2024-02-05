import { LitElement, html, css } from "https://cdn.jsdelivr.net/gh/lit/dist@3/core/lit-core.min.js"

export class RatingLines extends LitElement {
   static styles = css`
      .container {
         display: flex;
         gap: 0.5rem;
      }
      .line {
         border-width: 12px 0 12px 0;
         border-style: solid;
         border-color: transparent;
         height: 4px;
         flex-grow: 1;
         background-clip: padding-box;
      }
      .filled {
         background-color: white;
      }
      .unfilled {
         background-color: white;
         opacity: 0.3;
      }
      .pointer:hover {
         cursor: pointer;
         opacity: 0.7;
      }
   `

   static properties = {
      rating: { type: Number },
      maxrating: { type: Number },
      readonly: { type: Boolean },
   }

   constructor() {
      super()
      this.maxrating = 5
      this.rating = this.maxrating
   }

   setRating(rating) {
      if (this.readonly) return
      if (this.rating === 1 && rating === 1) rating = 0
      this.rating = rating
      //this.dispatchEvent(new CustomEvent("change", { detail: rating }))
   }

   render() {
      return html`
         <div class="container">
            ${Array.from(
               { length: this.rating },
               (_, i) =>
                  html`<span
                     class="line filled ${!this.readonly && "pointer"}"
                     @click=${() => this.setRating(i + 1)}
                  ></span>`
            )}
            ${Array.from(
               { length: this.maxrating - this.rating },
               (_, i) =>
                  html`<span
                     class="line unfilled ${!this.readonly && "pointer"}"
                     @click=${() => this.setRating(i + this.rating + 1)}
                  ></span>`
            )}
         </div>
      `
   }
}

customElements.define("rating-lines", RatingLines)
