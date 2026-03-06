# Football Match Predictor

A powerful, hybrid data-pipeline that automatically scrapes soccer match statistics and feeds them into a Deep Neural Network designed to predict match outcomes and betting targets.

## Overview
This repository contains a complete pipeline spanning automated data acquisition and predictive data science. It comprises two primary parts: a fast, concurrent Go routine (`godi.go`) that securely authenticates with a sports API to scrape, process, and compile an ongoing season's football records into a CSV format. Following the data collection, a Python-based Jupyter Notebook environment constructs a multi-layer deep neural network (using TensorFlow and Keras) to process the aggregated statistics and output binary predictions (e.g., predicting over/under thresholds like 'over 3.5' goals).

## Features
- **Golang API Scraper**: Automated HTTP POST sequences to interact with secure backend APIs.
- **Custom Sort Algorithms**: Implementation of quicksort in Go to properly align league tables by points and goal difference.
- **Deep MLP Classifier**: A Multi-Layer Perceptron (MLP) built in Keras with Sigmoid and ReLU activations.
- **End-to-End Pipeline**: From scraping raw JSON payloads to formatting training vectors for ML optimization.

## Tech Stack
- Python (Jupyter)
- Go (Golang)
- TensorFlow & Keras
- Pandas
- Scikit-Learn

## Project Architecture
```text
football-prediction-master/
  godi.go           # Golang utility for API scraping, JSON unmarshalling, and CSV generation
  football.ipynb    # Neural network initialization, dataset splitting, and model training
  data/             # Directories acting as localized storage for intermediate scraped files
```

## Installation
You'll need both the Go compiler and Python data science tools.

**For Data Scraping (Go):**
```bash
# Ensure Go is installed
go build godi.go
# Run the local executable to populate matches.csv
./godi
```

**For Prediction (Python):**
```bash
pip install tensorflow keras pandas numpy scikit-learn jupyter
```

## Running the Project
1. Run the `godi.go` script first to generate the latest required `.csv` files.
2. Launch the Jupyter Notebook environment to digest the collected data:
```bash
jupyter notebook football.ipynb
```

## Model Card

### Model Overview
The ML aspect is a feed-forward deep neural network (DNN) constructed to map 19 unique input variables into a single probabilistic outcome (value between 0 and 1).

### Model Architecture
- **Input Layer**: Flattened to handle tabular sequences (shape size 19).
- **Hidden Layers**: Dense neuron clusters of sizes `[64, 64, 32]` using mixed activations (`relu` and `tanh`) allowing the network to distinguish non-linear dependencies.
- **Output Layer**: A single node employing a `sigmoid` activation function to compute definitive percentage outcomes (e.g., predicting true or false for arbitrary target events like Total Goals > 3.5).

### Training Process
- Target column extracted from independent attributes.
- Random matrix cross-validation applied using an 95% Train / 5% Test split.
- Trained across 5000 epochs, tracking overall subset accuracy metrics and loss functions.

### Go Application Workflow
- Binds to live match endpoints (`odibets.com/api/`).
- Extracts matchweeks progressively down to matchweek 38 and recurses using built-in thread-sleeping.
- Parses 'GG' (Goal-Goal) and 'Over 1.5' accuracies for specialized analytics.

### Limitations
- Hardcoded API credentials and endpoints inside the Go script may require refactoring or environment variable handling for scaled deployments.
- Extended training schedules (5000 epochs on dense layers) risk overfitting without implemented callbacks or dropout mechanisms inside Keras.

## Professional Highlights
- **Demonstrated strong polyglot capabilities** integrating low-level, statically typed languages (Go) for high-performance scraping alongside high-level mathematical environments (Python/TensorFlow) for modeling.
- **Implemented custom Quicksort schemas** within Go contexts handling map slices to reconstruct league tables independent of external sorting packages.
- **Configured complex neural networks**, effectively designing dense multi-layer perceptrons capable of evaluating massive multivariate categorical dependencies inside sports analytics.

## License
MIT License

## Contributing
Contributions are welcome. Feel free to open issues or submit pull requests for enhancements.

## Author
Lih Ingabo
