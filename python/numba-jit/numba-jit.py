import click
import numba

from data import real_estate_array
from decorator import timing


@timing
def expmean(rea):
    """Regular Function"""

    val = rea.mean() ** 2
    return val


@timing
@numba.jit(nopython=True)
def expmean_jit(rea):
    """Perform multiple mean calculations"""

    val = rea.mean() ** 2
    return val


@click.command()
@click.option('--jit/--no-jit', default=False)
def jit_test(jit):
    rea = real_estate_array()
    if jit:
        click.echo(click.style('Running with JIT', fg='green'))
        expmean_jit(rea)
    else:
        click.echo(click.style('Running NO JIT', fg='red'))
        expmean(rea)


if __name__ == "__main__":
    jit_test()
