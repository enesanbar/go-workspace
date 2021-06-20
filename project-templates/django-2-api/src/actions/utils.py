from .models import Action


def create_action(user, verb, target=None):
    # no existing actions found
    action = Action(user=user, verb=verb, target=target)
    action.save()
