#!/bin/bash

# Define image names and tags
IMAGES=(
    "remahanrembulan/book-management-user:0.0.1"
    "remahanrembulan/book-management-book:0.0.1"
    "remahanrembulan/book-management-category:0.0.1"
    "remahanrembulan/book-management-author:0.0.1"
)

# Loop through the images and push each one
for IMAGE in "${IMAGES[@]}"; do
    echo "Pushing $IMAGE..."
    docker push "$IMAGE"
    if [ $? -ne 0 ]; then
        echo "Failed to push $IMAGE"
        exit 1
    fi
done

echo "All images pushed successfully."