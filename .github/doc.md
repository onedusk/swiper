 These GitHub Actions workflows are designed to automate tasks like code review and issue triage using Gemini. Here's
  how you can use them:

  How to Use

  These workflows are triggered by events in your GitHub repository.

   * Automatic Code Review: When you open a pull request, a code review will be automatically triggered.
   * Manual Code Review: To manually trigger a code review on a pull request, add a comment with @gemini-cli /review.
   * Automatic Issue Triage: When a new issue is created, it will be automatically triaged and labeled.
   * Manual Issue Triage: To manually trigger triage on an issue, add a comment with @gemini-cli /triage.
   * Scheduled Triage: Every hour, the system will look for any issues that haven't been triaged and label them.
   * General Purpose Commands: You can run a general command by commenting @gemini-cli <your command> on an issue or
     pull request.

  Configuration

  For these workflows to work, you need to configure secrets and variables in your repository's settings under
  Settings > Secrets and variables > Actions:

  Secrets:

   * GEMINI_API_KEY or GOOGLE_API_KEY: Your API key for the Gemini API.
   * APP_PRIVATE_KEY: The private key for your GitHub App (if you are using one for authentication).

  Variables:

   * APP_ID: The ID of your GitHub App.
   * GCP_WIF_PROVIDER: The Workload Identity Federation provider from Google Cloud.
   * GOOGLE_CLOUD_PROJECT: Your Google Cloud project ID.
   * GOOGLE_CLOUD_LOCATION: The Google Cloud location (e.g., us-central1).
   * SERVICE_ACCOUNT_EMAIL: The email of the Google Cloud service account.
   * GEMINI_MODEL: The Gemini model to use (e.g., gemini-1.5-pro-latest).
   * GOOGLE_GENAI_USE_VERTEXAI: Set to true to use Vertex AI.
   * GOOGLE_GENAI_USE_GCA: Set to true to use Gemini Code Assist.

  These workflows provide a powerful way to automate parts of your development lifecycle on GitHub.
