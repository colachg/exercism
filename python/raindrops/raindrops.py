def convert(n):
    drops = ((3, 'Pling'), (5, 'Plang'), (7, 'Plong'))
    """Converts a number to a string according to the raindrop sounds."""
    speak = [s for f, s in drops if n % f == 0]
    return "".join(speak) if speak else str(n)
