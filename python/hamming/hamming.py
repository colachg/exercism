def distance(strand_a, strand_b):
    """ distance is ..."""
    if len(strand_a) != len(strand_b):
        raise ValueError("invalid")
    return sum(i != j for i, j in zip(strand_a, strand_b))