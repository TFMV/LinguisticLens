gcloud auth configure-docker
docker build -t gcr.io/$PROJECT_ID/linguisticlens .
docker push gcr.io/$PROJECT_ID/linguisticlens

gcloud run deploy linguisticlens \
    --image gcr.io/$PROJECT_ID/linguisticlens \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated \
    --set-env-vars OPENAI_API_KEY=