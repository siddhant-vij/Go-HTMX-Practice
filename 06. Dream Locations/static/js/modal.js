function showConfirmationModal(event) {
  event.preventDefault();
  const action = event.detail.elt.dataset.action;
  const confirmationModal = `
      <dialog class="modal">
        <div id="confirmation">
          <h2>Are you sure?</h2>
          <p>Do you really want to ${action} this place?</p>
          <div id="confirmation-actions">
            <button id="action-no" className="button-text">
              No
            </button>
            <button id="action-yes" className="button">
              Yes
            </button>
          </div>
        </div>
      </dialog>
    `;
  document.body.insertAdjacentHTML("beforeend", confirmationModal);
  const dialog = document.querySelector("dialog");

  const noBtn = document.getElementById("action-no");
  noBtn.addEventListener("click", function () {
    dialog.remove();
  });

  const yesBtn = document.getElementById("action-yes");
  yesBtn.addEventListener("click", function () {
    event.detail.issueRequest();
    dialog.remove();
  });

  dialog.showModal();
}

// Only show this modal for dream & suggested locations
const dreamList = document.getElementById("dreamList");
dreamList.addEventListener("htmx:confirm", showConfirmationModal);

const suggestedList = document.getElementById("suggestedList");
suggestedList.addEventListener("htmx:confirm", showConfirmationModal);
