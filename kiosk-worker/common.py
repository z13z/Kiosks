import hashlib


def get_sha256(text):
    hash_generator = hashlib.sha256()
    hash_generator.update(text.encode('utf-8'))
    return hash_generator.hexdigest()
