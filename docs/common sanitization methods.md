# Common Sanitization Methods

> Here are some common sanitization methods you might find in cloud service environments:

- Data Deletion: This is the most basic method and involves deleting data from a cloud storage resource. Deleted data can often be recovered within a certain retention period unless additional steps are taken.
- Data Overwrite: Some cloud providers offer data overwrite methods that involve writing random or null data over the existing data, making it more challenging to recover the original content.
- Encryption Key Rotation: For encrypted data, rotating encryption keys can effectively make data inaccessible. When the encryption key is changed, the data encrypted with the old key becomes unreadable.
- Data Obfuscation: Data obfuscation involves altering the data in a way that it is still available but not interpretable. This can include techniques like data masking or tokenization.
- Secure Delete: Secure delete methods go beyond basic deletion by overwriting the data with random values multiple times, making it extremely difficult to recover.
- Cryptographic Erasure: Some cloud providers offer cryptographic erasure methods that involve the use of cryptographic techniques to make data irretrievable. This may include encrypting the data with a secure key that is then destroyed.
- Logical Deletion: Logical deletion refers to removing references to data, making it appear as if it has been deleted, even though it might still exist in the storage. Over time, the storage system may reclaim this space.
