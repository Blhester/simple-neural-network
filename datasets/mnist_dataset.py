import os

from datasets import load_dataset

# print('Downloaded dataset... Saving to file... tmp/mnist...')

# dataset.save_to_disk('tmp/mnist')
csv_file_path = '../csv_reader/output.csv'  # Specify the desired output CSV file path
output_directory = 'tmp/images'  # Specify the desired output directory

counter = 0
if __name__ == "__main__":
    dataset = load_dataset('mnist', split='train')
    for idx, data in enumerate(dataset):
        counter += 1
        # Extract the image object and the label
        image, label = data['image'], data['label']

        expected_directory = os.path.join(output_directory, f"{int(label)}")
        if not os.path.exists(expected_directory):
            os.makedirs(expected_directory)
        # Generate the output file path
        # Naming pattern: "label_index.png" (e.g., "5_0.png" for the first image labeled as 5)
        output_path = os.path.join(expected_directory, f"{counter}.png")

        # Save the image
        image.save(output_path)
    print("Images saved successfully.")
