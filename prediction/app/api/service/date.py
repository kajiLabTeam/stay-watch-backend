def time_to_minuts(time: str) -> int:
    """Convert time string to minutes"""
    h, m = map(int, time.split(":"))
    return h * 60 + m
